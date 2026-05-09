package url

import "testing"

var parseTests = []struct {
	name string
	uri  string
	want *URL
}{
	{
		name: "full",
		uri:  "https://meleu.sh/shellcheck",
		want: &URL{
			Scheme: "https",
			Host:   "meleu.sh",
			Path:   "shellcheck",
		},
	}, {
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
		t.Logf("run: %s", tt.name)

		got, err := Parse(tt.uri)
		if err != nil {
			t.Fatalf("Parse(%q) err = %q, want <nil>", tt.uri, err)
		}
		if *got != *tt.want {
			t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", tt.uri, got, tt.want)
		}
	}
}

func TestURLString(t *testing.T) {
	u := &URL{
		Scheme: "https",
		Host:   "meleu.sh",
		Path:   "shellcheck",
	}

	got := u.String()
	want := "https://meleu.sh/shellcheck"
	if got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}
