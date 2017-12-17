package css

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

var SYMBOL_TABLE map[rune]bool

var SYMBOLS []rune = []rune{
	'*', '$', '!', '-', '~', ',', '#',
	'.', '>', '^', '=', '+', ':', ';',
	'(', ')',
	'{', '}',
}

func init() {
	SYMBOL_TABLE = map[rune]bool{}
	for _, sym := range SYMBOLS {
		SYMBOL_TABLE[sym] = true
	}
}

type CssLexer struct {
	input string
	pos   uint
}

func IsIdentifier(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-'
}

func IsNumber(r rune) bool {
	return unicode.IsDigit(r)
}

func TokenizeCss(input string) []*Token {
	c := &CssLexer{
		input: input,
		pos:   0,
	}
	var tokens []*Token
	for c.hasNext() {
		if token := c.recognizeToken(); token != nil {
			tokens = append(tokens, token)
		}
	}
	return tokens
}

func (c *CssLexer) recogniseSymbol() *Token {
	sym, _ := c.consume()
	return NewToken([]rune{sym}, Symbol)
}

func (c *CssLexer) recognizeString() *Token {
	fst, _ := c.consume()
	lexeme := c.consumeWhile(func(r rune) bool {
		return r != fst
	})
	c.expect(fst)

	lexeme = append([]rune{fst}, lexeme...)
	lexeme = append(lexeme, fst)
	return NewToken(lexeme, String)
}

func (c *CssLexer) recognizeToken() *Token {
	c.skipLayout()
	curr := c.peek(0)

	// clean this up some point!
	if unicode.IsLetter(curr) || curr == '-' && (c.hasNext() && (c.peek(1) == '_' || unicode.IsLetter(c.peek(1)))) || curr == '_' {
		// identifiers must start with a letter, or an underscore
		// they can start with a hyphen BUT this must be followed
		// by a letter or an underscore.
		return c.recognizeIdentifier()
	} else if unicode.IsDigit(curr) || (curr == '.' && (c.hasNext() && IsNumber(c.peek(1)))) {
		return c.recognizeNumber()
	} else if curr == '/' {
		c.skipComment()
		return nil
	} else if curr == '"' || curr == '\'' {
		return c.recognizeString()
	}

	if _, ok := SYMBOL_TABLE[curr]; ok {
		return c.recogniseSymbol()
	}

	fmt.Println("oh dear!", string(curr), "is unhandled! next char is'"+string(c.peek(1))+"'")
	return nil
}

func (c *CssLexer) recognizeNumber() *Token {
	startsWithDot := c.peek(0) == '.'
	if startsWithDot {
		c.consume()
	}
	lexeme := c.consumeWhile(IsNumber)
	if startsWithDot {
		lexeme = append([]rune{'.'}, lexeme...)
	} else if c.hasNext() {
		// this is going by the assumption that if it
		// doesnt start with a dot, then technically
		// we will check if it is a floating number
		// if we didnt put this in the else then the lexer
		// would allow .3255.5444
		if c.peek(0) == '.' {
			c.consume()

			var precisionyPart []rune
			precisionyPart = append(precisionyPart, '.')
			precisionyPart = append(precisionyPart, c.consumeWhile(IsNumber)...)
			lexeme = append(lexeme, precisionyPart...)
		}
	}
	return NewToken(lexeme, Number)
}

func (c *CssLexer) recognizeIdentifier() *Token {
	value := c.consumeWhile(IsIdentifier)
	tokenType := Identifier
	// we can handle keywords
	// here but this might be a bit
	// complicated
	return NewToken(value, tokenType)
}

func (c *CssLexer) skipLayout() {
	c.consumeWhile(func(r rune) bool {
		// layout is spaces, tabs, etc.
		// in ascii this is anything that
		// is below the space, or 32
		return r <= ' '
	})
}

// these functions have literally been copied and pasted
// from the HtmlParser. I feel like we could have a module
// that covers this but its use is so small i dont feel it
// justifies it just yet...
func (c *CssLexer) hasNext() bool {
	return c.pos < uint(len(c.input))
}

func (c *CssLexer) consumeWhile(predicate func(rune) bool) []rune {
	// we can either slice here or we
	// can append it to a temporary buffer
	// thing. for now we will append it to
	// a buffer because it's a lot easier!

	buffer := []rune{}
	for c.hasNext() && predicate(c.peek(0)) {
		value, _ := c.consume()
		if value == utf8.RuneError {
			panic("oh dear!")
			break
		}
		buffer = append(buffer, value)
	}
	return buffer
}

func (c *CssLexer) expect(r ...rune) {
	for _, expected := range r {
		if c.hasNext() && c.peek(0) != expected {
			var offs uint = 10
			sample := c.input[c.pos : c.pos+offs]
			panic("expected '" + string(r) + "', got '" + string(sample) + "'")
		}
		c.consume()
	}
}

func (c *CssLexer) consume() (rune, uint) {
	r, signedSize := utf8.DecodeRuneInString(c.input[c.pos:])
	if signedSize < 0 {
		panic("shit")
	}
	size := uint(signedSize)
	c.pos += size
	return r, size
}

func (c *CssLexer) peek(offs uint) rune {
	r, _ := utf8.DecodeRuneInString(c.input[c.pos+offs:])
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
