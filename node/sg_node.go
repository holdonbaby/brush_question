package node

type SNode struct {
	Son      *Node
	ValueStr string
	ValueInt int
}

func NewSNode() *SNode {
	return &SNode{}
}

func NewSNodeWithContent(son *Node, valueStr string, valueInt int) *SNode {
	return &SNode{
		Son:      son,
		ValueStr: valueStr,
		ValueInt: valueInt,
	}
}
