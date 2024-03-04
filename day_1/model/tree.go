package model

type Node struct {
	Value     string
	LeftNode  *Node
	RightNode *Node
}

func NewNode(value string) *Node {
	return &Node{
		Value: value,
	}
}

func NewBinaryExpressionTree(rootValue string, left *Node, right *Node) *Node {
	return &Node{
		Value:     rootValue,
		LeftNode:  left,
		RightNode: right,
	}
}
