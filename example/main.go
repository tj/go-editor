package main

import (
	"fmt"
	"log"

	"github.com/tj/go-editor"
)

func main() {
	b, err := editor.Read()
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	fmt.Printf("%q\n", string(b))
}
