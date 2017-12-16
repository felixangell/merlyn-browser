package dom

type TextNode struct {
	SimpleDomNode
	value string
}

func NewTextNode(value string) *TextNode {
	return &TextNode{
		value: value,
	}
}