package main

import (
	"bytes"
	"fmt"

	"github.com/github-tijlxyz/goldmark-nostr/package/extension"
	"github.com/yuin/goldmark"
)

func main() {
	somecontent := ""

	md := goldmark.New(goldmark.WithExtensions(extension.NewNostr()))

	var buf bytes.Buffer
	if err := md.Convert([]byte(somecontent), &buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
