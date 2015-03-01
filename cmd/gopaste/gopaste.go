package main

import (
	"github.com/elimisteve/clipboard"

	"fmt"
)

func main() {
	data, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))
}
