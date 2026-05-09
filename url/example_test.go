package url_test

import (
	"fmt"
	"log"

	"github.com/meleu/gbe-book/url"
)

func ExampleParse() {
	uri, err := url.Parse("https://meleu.sh/shellcheck")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uri)
	// Output:
	// https://meleu.sh/shellcheck
}
