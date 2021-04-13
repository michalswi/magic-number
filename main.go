package main

import (
	"fmt"
	"log"

	"github.com/michalswi/magic-numbers/magic"
)

func main() {
	magic, err := magic.MagicNumbers("/bin/bash")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(magic)
}
