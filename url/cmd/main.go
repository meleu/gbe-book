package main

import (
	"fmt"
	"log"

	"github.com/meleu/gbe-book/url"
)

func main() {
	uri, err := url.Parse("https://meleu.sh/shellcheck")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", uri.Scheme)
	fmt.Println("Host..:", uri.Host)
	fmt.Println("Path..:", uri.Path)
	fmt.Println(uri)
}
