package dom

type NodeType int

const (
	Text NodeType = iota
	Element
)

type DomNode interface {
	GetChildren() []DomNode
	GetType() NodeType
}

type SimpleDomNode struct {
	children []DomNode
	kind NodeType
}

func (s *SimpleDomNode) GetChildren() []DomNode {
	return s.children
}

func (s *SimpleDomNode) GetType() NodeType {
	return s.kind
}