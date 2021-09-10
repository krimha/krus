package krus

import "testing"

func TestNodeSet(t *testing.T) {

	nodes := make([]*node, 6)
	for i, name := range []string{"q1", "q2", "q3", "q4"} {
		nodes[i] = newNode(name)
	}

	setA := newNodeSet()
	setA.Insert(nodes[0])
	setA.Insert(nodes[1])
	setA.Insert(nodes[2])
	
	setB := newNodeSet()
	setB.Insert(nodes[1])
	setB.Insert(nodes[2])
	setB.Insert(nodes[3])

	setA.InsertSet(setB)

	result := len(setA.storage)
	if result != 4 {
		t.Fatalf("setA contains the wrong number of elements %d", result)
	}

	if !setA.Contains(nodes[3]) {
		t.Fatalf("setA does not contain a node from setB")
	}
}

func TestSize(t *testing.T) {
	setA := newNodeSet()
	n := newNode("q")
	setA.Insert(n)
	setA.Insert(n)
	
	if setA.Size() != 1 {
		t.Fatalf("Size is wrong")
	}
}