package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func readFile(path string, content *string) error {
	f, e := os.Open(path)
	if nil != e {
		return e
	}
	defer f.Close()

	buf, e := ioutil.ReadAll(f)
	if nil != e {
		return e
	}

	*content = string(buf)
	return nil
}

func selectHtml(htmlContent, selector, split string) string {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(htmlContent)))
	if err != nil {
		panic(fmt.Errorf("wrong html format %v", err))
	}

	sel := doc.Find(selector)

	var buf bytes.Buffer

	for _, n := range sel.Nodes {
		buf.Write([]byte("\n" + split + "\n"))
		if err := html.Render(&buf, n); nil != err {
			panic(fmt.Errorf("rend node fail %v", err))
		}
	}
	return buf.String()
}

func main() {
	var selector string
	var content string
	var contentFile string
	var split string
	flag.StringVar(&selector, "selector", ".container", "xpath selector")
	flag.StringVar(&content, "content", "", "html text content")
	flag.StringVar(&contentFile, "contentFile", "./1.html", "html text content file")
	flag.StringVar(&split, "split", "-----", "split line content of each matched selection")
	flag.Parse()

	if len(content) <= 0 && len(contentFile) > 0 {
		if err := readFile(contentFile, &content); nil != err {
			panic("read file fail" + err.Error())
		}
	}

	if len(selector) <= 0 || len(content) <= 0 {
		fmt.Println(flag.ErrHelp)
		panic("wrong parameter")
	}

	fmt.Println(selectHtml(content, selector, split))
}
