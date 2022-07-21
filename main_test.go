package main

import (
	"fmt"
	"testing"
)

func Test_SelectClass(t *testing.T) {
	var content string
	if err := readFile("./1.html", &content); nil != err {
		t.Error(err)
	}
	fmt.Println(selectHtml(content, ".container", "-----"))
}

func Test_SelectID(t *testing.T) {
	var content string
	if err := readFile("./1.html", &content); nil != err {
		t.Error(err)
	}
	fmt.Println(selectHtml(content, "#header", "-----"))
}

func Test_SelectTag(t *testing.T) {
	var content string
	if err := readFile("./1.html", &content); nil != err {
		t.Error(err)
	}
	fmt.Println(selectHtml(content, "img", "-----"))
}

func Test_SelectNest(t *testing.T) {
	var content string
	if err := readFile("./1.html", &content); nil != err {
		t.Error(err)
	}
	fmt.Println(selectHtml(content, "head>meta", "-----"))
}

func Test_SelectNest1(t *testing.T) {
	var content string
	if err := readFile("./1.html", &content); nil != err {
		t.Error(err)
	}
	fmt.Println(selectHtml(content, ".container>.navbar-brand", "-----"))
}

func Test_SelectNest2(t *testing.T) {
	var content string
	if err := readFile("./1.html", &content); nil != err {
		t.Error(err)
	}
	fmt.Println(selectHtml(content, ".container .navbar-brand", "-----"))
}
