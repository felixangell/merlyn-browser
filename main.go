package main

import (
	"fmt"
	"github.com/felixangell/merlyn/html"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, World!")
	
	fileBytes, err := ioutil.ReadFile("tests/simple_page.html")
    if err != nil {
        fmt.Print(err)
    }
    
    html.ParseHtml(string(fileBytes))
}
