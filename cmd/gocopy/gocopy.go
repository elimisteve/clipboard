package main

import (
	"github.com/elimisteve/clipboard"

	"io/ioutil"
	"os"
)

func main() {
	out, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	if err := clipboard.WriteAll(out); err != nil {
		panic(err)
	}
}
