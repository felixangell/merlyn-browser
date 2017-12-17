package css

import (
	_ "fmt"
	"unicode/utf8"
)

type CssLexer struct {
	input string
	pos   uint
}

func TokenizeCss(input string) []*Token {
	c := &CssLexer{
		input: input,
		pos:   0,
	}

	var tokens []*Token
	for c.hasNext() {

	}
	return tokens
}

// these functions have literally been copied and pasted
// from the HtmlParser. I feel like we could have a module
// that covers this but its use is so small i dont feel it
// justifies it just yet...
func (p *CssLexer) hasNext() bool {
	return p.pos < uint(len(p.input))
}

func (p *CssLexer) consumeWhile(predicate func(rune) bool) []rune {
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

func (p *CssLexer) expect(r ...rune) {
	for _, expected := range r {
		if p.hasNext() && p.peek(0) != expected {
			var offs uint = 10
			sample := p.input[p.pos : p.pos+offs]
			panic("expected '" + string(r) + "', got '" + string(sample) + "'")
		}
		p.consume()
	}
}

func (p *CssLexer) consume() (rune, uint) {
	r, signedSize := utf8.DecodeRuneInString(p.input[p.pos:])
	if signedSize < 0 {
		panic("shit")
	}
	size := uint(signedSize)
	p.pos += size
	return r, size
}

func (p *CssLexer) peek(offs uint) rune {
	r, _ := utf8.DecodeRuneInString(p.input[p.pos:])
	if r == utf8.RuneError {
		panic("oh dear!")
	}
	return r
}

func (c *CssLexer) skipComment() {
	c.expect('/', '*')
	c.consumeWhile(func(r rune) bool {
		return r != '*'
	})
	c.expect('*', '/')
}
