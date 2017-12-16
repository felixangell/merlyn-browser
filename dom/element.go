package dom

type AttributeMap = map[string]string

type DomElement struct {
	SimpleDomNode
	attributes AttributeMap
	name string
}

func NewDomElement(name string) *DomElement {
	return &DomElement {
		attributes: AttributeMap{},
		name: name,
	}
}