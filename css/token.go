package css

type TokenType uint

const (
	Identifier TokenType = iota
	Number
	String
	Keyword
	Symbol
)

type Token struct {
	lexeme string
	kind   TokenType
}

func NewToken(lexeme string, kind TokenType) *Token {
	return &Token{
		lexeme: lexeme,
		kind:   kind,
	}
}
