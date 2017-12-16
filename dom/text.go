package dom

import "fmt"

type TextNode struct {
	SimpleDomNode
	value string
}

func NewTextNode(value string) *TextNode {
	return &TextNode{
		value: value,
	}
}

func (t *TextNode) Info() string {
	return fmt.Sprintf("t %s", t.value)
}
