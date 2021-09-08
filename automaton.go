package krus

type node struct {
	name string
	edges map[byte]*node
}

func NewNode(name string) *node {
	return &node { name, make(map[byte]*node) }
}