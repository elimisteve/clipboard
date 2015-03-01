package clipboard_test

import (
	"github.com/elimisteve/clipboard"

	"fmt"
)

func Example() {

	clipboard.WriteAll([]byte("日本語"))
	text, _ := clipboard.ReadAll()
	fmt.Println(string(text))

	// Output:
	// 日本語
}
