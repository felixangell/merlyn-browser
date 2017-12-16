package dom

type AttributeMap = map[string]string

type ElementNode struct {
	SimpleDomNode
	attributes AttributeMap
	name string
}

func NewElementNode(name string, children []Node) *ElementNode {
	return &ElementNode {
		SimpleDomNode: SimpleDomNode {
			kind: Element,
			children: children,
		},
		attributes: AttributeMap{},
		name: name,
	}
}