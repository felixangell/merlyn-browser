package css

type CssParser struct {
	input []*Token
	pos   uint
}

func (c *CssParser) parseRule() *Rule {

	return nil
}

func (c *CssParser) parseRules() []*Rule {
	var rules []*Rule
	return rules
}

func ParseCss(tokens []*Token) *StyleSheet {
	c := &CssParser{
		input: tokens,
	}
	return &StyleSheet{
		rules: c.parseRules(),
	}
}
