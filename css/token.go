//go:generate stringer -type=TokenType

package css

type TokenType uint

const (
	Identifier TokenType = iota
	Number
	String
	Symbol
)

type Token struct {
	lexeme []rune
	kind   TokenType
}

func NewToken(lexeme []rune, kind TokenType) *Token {
	return &Token{
		lexeme: lexeme,
		kind:   kind,
	}
}
