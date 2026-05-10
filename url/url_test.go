package url

import (
	"fmt"
	"strings"
	"testing"
)

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{
		name: "opaque URL",
		uri:  "data:text/plain",
		want: &URL{Scheme: "data"},
	},
	{
		name: "with data scheme",
		uri:  "data:text/plain;base64,R28gYnkgRXhhbXBsZQ==",
		want: &URL{Scheme: "data"},
	},
	{
		name: "full",
		uri:  "https://meleu.sh/shellcheck",
		want: &URL{
			Scheme: "https",
			Host:   "meleu.sh",
			Path:   "shellcheck",
		},
	},
	{
		name: "without path",
		uri:  "https://meleu.sh",
		want: &URL{
			Scheme: "https",
			Host:   "meleu.sh",
			Path:   "",
		},
	},
}

func TestParseTable(t *testing.T) {
	for _, tt := range parseTests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.uri)
			if err != nil {
				t.Fatalf("Parse(%q) err = %q, want <nil>", tt.uri, err)
			}
			if *got != *tt.want {
				t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
			}
		})
	}
}

func TestParseError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{name: "without scheme", uri: "github.com"},
		{name: "empty scheme", uri: "://github.com"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.uri)
			if err == nil {
				t.Errorf("Parse(%q) err=nil; want an error", tt.uri)
			}
		})
	}
}

func TestURLString(t *testing.T) {
	tests := []struct {
		name string
		uri  *URL
		want string
	}{
		{
			name: "nil",
			uri:  nil,
			want: "",
		},
		{
			name: "empty",
			uri:  new(URL),
			want: "",
		},
	}

	for _, tt := range tests {
		got := tt.uri.String()
		if got != tt.want {
			t.Errorf("\ngot  %q\nwant %q\nfor  %#v", got, tt.want, tt.uri)
		}
	}
}

func BenchmarkURLString(b *testing.B) {
	u := &URL{
		Scheme: "https",
		Host:   "meleu.sh",
		Path:   "shellcheck",
	}

	for b.Loop() {
		_ = u.String()
	}
}

func BenchmarkURLStringLong(b *testing.B) {
	for _, n := range []int{10, 100, 1_000} {
		u := &URL{
			Scheme: strings.Repeat("x", n),
			Host:   strings.Repeat("y", n),
			Path:   strings.Repeat("z", n),
		}

		b.Run(fmt.Sprintf("%d", n), func(b *testing.B) {
			for b.Loop() {
				_ = u.String()
			}
		})
	}
}
