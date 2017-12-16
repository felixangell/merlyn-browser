package dom

import "fmt"

const TAB_SIZE = 2

func printNodes(nodes []Node, level int) {
	paddingRaw := []byte{'>'}
	for idx := 1; idx < level*TAB_SIZE; idx++ {
		paddingRaw = append(paddingRaw, ' ')
	}
	padding := string(paddingRaw)

	for _, node := range nodes {
		switch n := node.(type) {
		case *TextNode:
			fmt.Println(padding, n.Info())
		case *ElementNode:
			fmt.Println(padding, n.Info())
			printNodes(n.children, level+1)
		default:
			fmt.Println("what is this?")
		}
	}
}

func TreePrinter(nodes []Node) {
	printNodes(nodes, 0)
}
