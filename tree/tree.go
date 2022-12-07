package tree

type Node struct {
	name     string
	size     int
	children []*Node
	parent   *Node
}

func NewNode(name string) *Node {
	return &Node{name: name}
}

func (n *Node) addChild(node *Node) {
	n.children = append(n.children, node)
}
