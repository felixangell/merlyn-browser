package html

import (
	"fmt"
	"github.com/felixangell/merlyn/dom"
	"unicode/utf8"
)

type HtmlParser struct {
	input string
	pos uint
}

func ParseHtml(htmlCode string) []dom.DomNode {
	p := &HtmlParser{
		input: htmlCode,
	}

	var nodes []dom.DomNode
	for p.hasNext() {
		fmt.Println(string(p.consume()))
	}
	return nodes
}

func (p *HtmlParser) hasNext() bool {
	return p.pos < uint(len(p.input))
}

func (p *HtmlParser) consume() rune {
	r, size := utf8.DecodeRuneInString(p.input[p.pos:])
	p.pos += uint(size)
	return r
}

func (p *HtmlParser) peek(offs uint) rune {
	r, _ := utf8.DecodeRuneInString(p.input[p.pos:])
	if r == utf8.RuneError {
		panic("oh dear!")
	}
	return r
}

func (p *HtmlParser) parseNode() dom.DomNode {
	return nil
}