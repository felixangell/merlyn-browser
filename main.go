package main

import (
	"fmt"
	"github.com/felixangell/merlyn/css"
	"github.com/felixangell/merlyn/dom"
	"github.com/felixangell/merlyn/html"
	"io/ioutil"
)

func readToString(path string) string {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	return string(fileBytes)
}

func main() {
	fmt.Println("Hello, World!")

	htmlDOM := html.ParseHtml(readToString("tests/simple_page.html"))
	dom.TreePrinter(htmlDOM)

	cssCode := readToString("tests/base.css")
	tokens := css.TokenizeCss(cssCode)
	css.TokenPrinter(tokens)
	css.ParseCss(tokens)
}
