package css

// this is me guessing the CSS grammar in EBNF
// probably a lot more complex however:
//
// Unit = "em" | "px" | "pt" | "rem" | "%" | ... etc
// Value = Identifier | Number [ Unit ];
// String_Literal = '"' { ... } '"'
// Operator = '*' | '$' | '^' | '~'
// Attribute_Set = "[" Identifier [Operator] "=" String_Literal "]"
// Psuedo_Class = ":" Identifier [ "(" [ Value ] { Value } ")" ]
// Selector = Identifier [ Psuedo_Class | Attribute_Set ];
// Combinator = '>' | '+' | '~' | '#'
// Selection = Selector { [ Combinator ] Selector }
// Rule = Selection Rule_Body
// Rule_Body = "{" { Declaration } "}"
// Declaration = Identifier ":" Value ";"
//

// hm?
// div.error
// a.error

// h1
// h1:first-child
// h1:nth-child(5)
// input[type="text"]
type Selector struct {
	name          string              // h1
	psuedoClasses map[string][]string // key => "first-child", value => [ "5", ]
	attributes    map[string]string   // key => "type", value = "text"
}

type Rule struct {
}
