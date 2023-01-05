//go:build ignore

package main

import (
	"fmt"
	translator "github.com/solywsh/go-googletrans"
)

var content = `你好，世界！`

func main() {
	t := translator.New(
		translator.WithProxy("http://127.0.0.1:7890"),
		translator.WithUserAgent("Custom Agent"),
		translator.WithServiceUrl("translate.google.com.hk"),
	)
	result, err := t.Translate(content, "auto", "en")
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Text)

	// Latency test
	t.Latency(t.AllServiceUrls())
}
