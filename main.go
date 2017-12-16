package main

import (
	"fmt"
	"github.com/felixangell/merlyn/html"
	"github.com/felixangell/merlyn/dom"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, World!")
	
	fileBytes, err := ioutil.ReadFile("tests/simple_page.html")
    if err != nil {
		fmt.Print(err)
    }
    
    htmlDOM := html.ParseHtml(string(fileBytes))
	dom.TreePrinter(htmlDOM)
}