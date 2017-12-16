package dom

type DomText struct {
	SimpleDomNode
	value string
}

func NewTextNode(value string) *DomText {
	return &DomText{
		value: value,
	}
}