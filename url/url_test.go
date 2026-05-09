package url

import "testing"

func TestParse(t *testing.T) {
	const uri = "https://meleu.sh/shellcheck"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want <nil>", uri, err)
	}

	want := &URL{
		Scheme: "https",
		Host:   "meleu.sh",
		Path:   "shellcheck",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
	}
}

func TestParseWithoutPath(t *testing.T) {
	const uri = "https://meleu.sh"

	got, err := Parse(uri)
	if err != nil {
		t.Fatalf("Parse (%q) err = %q, want <nil>", uri, err)
	}

	want := &URL{
		Scheme: "https",
		Host:   "meleu.sh",
		Path:   "",
	}
	if *got != *want {
		t.Errorf("Parse(%q)\ngot  %#v\nwant %#v", uri, got, want)
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
