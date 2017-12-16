package html

import (
	"fmt"
	"github.com/felixangell/merlyn/dom"
	"unicode/utf8"
	"strings"
	"unicode"
)

type HtmlParser struct {
	input string
	pos uint
}

func ParseHtml(htmlCode string) []dom.Node {
	p := &HtmlParser{
		input: htmlCode,
	}
	fmt.Println("Parsing some html code, yay!")
	return p.parseNodes()
}

func (p *HtmlParser) parseNodes() []dom.Node {
	var nodes []dom.Node
	for {
		p.consumeWhile(func (r rune) bool {
			return r <= ' '	
		})

		if !p.hasNext() || strings.HasPrefix(p.input[p.pos:], "</") {
			break
		}

		if node := p.parseNode(); node != nil {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (p *HtmlParser) hasNext() bool {
	return p.pos < uint(len(p.input))
}

func (p *HtmlParser) consumeWhile(predicate func(rune) bool) []rune {
	// we can either slice here or we 
	// can append it to a temporary buffer
	// thing. for now we will append it to
	// a buffer because it's a lot easier!

	buffer := []rune{}
	for p.hasNext() && predicate(p.peek(0)) {
		value, _ := p.consume()
		if value == utf8.RuneError {
			panic("oh dear!")
			break
		}
		buffer = append(buffer, value)
	}
	return buffer
}

func (p *HtmlParser) expect(r ...rune) {
	for _, expected := range r {
		if p.hasNext() && p.peek(0) != expected {
			var offs uint = 10
			sample := p.input[p.pos:p.pos + offs]
			panic("expected '" + string(r) + "', got '" + string(sample) + "'")
		}
		p.consume()
	}
}

func (p *HtmlParser) consume() (rune, uint) {
	r, signedSize := utf8.DecodeRuneInString(p.input[p.pos:])
	if signedSize < 0 {
		panic("shit")
	}
	size := uint(signedSize)
	p.pos += size
	return r, size
}

func (p *HtmlParser) peek(offs uint) rune {
	r, _ := utf8.DecodeRuneInString(p.input[p.pos:])
	if r == utf8.RuneError {
		panic("oh dear!")
	}
	return r
}

func IsHtmlTagRune(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func (p *HtmlParser) parseString() string {
	if p.peek(0) != '"' {
		return ""
	}
	p.expect('"')
	value := p.consumeWhile(func(r rune) bool { return r != '"' })
	p.expect('"')

	const quote string = "\""
	return quote + string(value) + quote
}

func (p *HtmlParser) parseElement() *dom.ElementNode {
	p.expect('<')
	name := p.consumeWhile(IsHtmlTagRune)
	
	attribs := dom.AttributeMap{}
	for {
		p.consumeWhile(func(r rune) bool { return r <= ' ' })
		if p.peek(0) == '>' {
			break
		}

		// this means that 
		// <div this is my thingy    ="foo">
		// would be a valid attribute...
		// stored as "this is my thingy    "
		// should we allow this?

		attributeName := p.consumeWhile(func (r rune) bool {
			return r != '='
		})
		p.expect('=')
		attributeValue := p.parseString()
		attribs[string(attributeName)] = attributeValue
	}

	p.expect('>')
	
	children := p.parseNodes()

	p.expect('<')
	p.expect('/')
	p.expect(name...)
	p.expect('>')

	return dom.NewElementNode(string(name), children, attribs)
}

func (p *HtmlParser) parseText() *dom.TextNode {
	value := p.consumeWhile(func(r rune) bool {
		return r != '<'		
	})
	return dom.NewTextNode(string(value))
}

func (p *HtmlParser) parseNode() dom.Node {
	switch p.peek(0) {
	case '<':
		return p.parseElement()
	default:
		return p.parseText()
	}
	return nil
}