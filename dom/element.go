package dom

import "fmt"

type AttributeMap map[string]string

func (a AttributeMap) Info() string {
	var result string
	idx := 0
	for key, val := range a {
		if idx > 0 {
			result += " "
		}
		result += fmt.Sprintf("%s=%s", key, val)
		idx++
	}
	return result
}

type ElementNode struct {
	SimpleDomNode
	attributes AttributeMap
	name       string
}

func NewElementNode(name string, attributes AttributeMap) *ElementNode {
	return &ElementNode{
		SimpleDomNode: SimpleDomNode{
			kind:     Element,
			children: []Node{},
		},
		attributes: attributes,
		name:       name,
	}
}

func (e *ElementNode) Info() string {
	return fmt.Sprintf("e %s, %d: [%s]", e.name, len(e.children), e.attributes.Info())
}
