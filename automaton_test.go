package krus

import "testing"

func TestNewNode(t *testing.T) {
	n := newNode("q1")
	result := n.name
	expected := "q1"
	if result != expected {
		t.Fatalf(`NewNode("q1").name = %s, exptected %s`, result, expected)
	}
}

func TestNewGraph(t *testing.T) {
	g := NewGraph([]string{"q0", "q1"}, "q0", []string{"q1"})

	if g.nodes["q0"] == nil {
		t.Fatalf("Node q0 was not initializeed")
	}

	if g.nodes["q0"].name != "q0" {
		t.Fatalf(`Node q0 was given wrong name "%s"`, g.nodes["q0"].name)
	}

	if g.start != g.nodes["q0"] {
		t.Fatalf("Start node was not set correctly")
	}
}

func TestConnect(t *testing.T) {
	g := NewGraph([]string{"q0", "q1"}, "q0", []string{"q1"})
	g.Connect("q0", "q1", '1')

	q0 := g.nodes["q0"]
	q1 := g.nodes["q1"]
	if !q0.edges['1'].Contains(q1) {
		t.Fatalf("Connect does not work")
	}

}

func TestMatch(t *testing.T) {
	g := NewGraph([]string{"q0", "q1"}, "q0", []string{"q1"})
	g.Connect("q0", "q1", '1')
	g.Connect("q1", "q0", '0')
	g.Connect("q0", "q0", '0')
	g.Connect("q1", "q1", '1')

	if g.Match("00000") {
		t.Fatalf(`Should not match "00000"`)
	}
	if !g.Match("000111") {
		t.Fatalf(`Should match "000111"`)
	}

}

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