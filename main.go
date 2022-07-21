package main

import (
	"bytes"
	"flag"
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	var selector string
	var content string
	var split string
	flag.StringVar(&selector, "selector", "", "xpath selector")
	flag.StringVar(&content, "content", "", "html text content")
	flag.StringVar(&split, "split", "-----", "split line content of each matched selection")
	flag.Parse()

	if len(selector) <= 0 || len(content) <= 0 {
		fmt.Println(flag.ErrHelp)
		panic("wrong parameter")
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(content)))
	if err != nil {
		panic(fmt.Errorf("wrong html format %v", err))
	}

	sel := doc.Find(selector)

	sel.Each(func(i int, nsel *goquery.Selection) {
		fmt.Printf("%v\n%v", split, nsel.Text())
	})
}
