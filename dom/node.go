package dom

type NodeType int

const (
	Text NodeType = iota
	Element
)

type Node interface {
	GetChildren() []Node
	GetType() NodeType
	Info() string
}

type SimpleDomNode struct {
	children []Node
	kind NodeType
}

func (s *SimpleDomNode) GetChildren() []Node {
	return s.children
}

func (s *SimpleDomNode) GetType() NodeType {
	return s.kind
}

func (s *SimpleDomNode) SetChildren(children []Node) {
	s.children = children
}

func (s *SimpleDomNode) SetKind(kind NodeType) {
	s.kind = kind
}