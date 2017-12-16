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