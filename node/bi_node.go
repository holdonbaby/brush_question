package node

type Node struct {
	Left     *Node
	Right    *Node
	ValueStr string
	ValueInt int
}

func NewNode() *Node {
	return &Node{}
}

func NewNodeWithContent(left, right *Node, valueStr string, valueInt int) *Node {
	return &Node{
		Left:     left,
		Right:    right,
		ValueStr: valueStr,
		ValueInt: valueInt,
	}
}
