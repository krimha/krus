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

func TestNewNonDeterministic(t *testing.T) {
	g := NewGraph([]string{"X", "0", "1", "2", "3"}, "X", []string{"3"})
	g.Connect("X", "X", '0')
	g.Connect("X", "X", '1')
	g.Connect("X", "0", '1')
	g.Connect("0", "1", '0')
	g.Connect("0", "1", '1')
	g.Connect("1", "2", '0')
	g.Connect("1", "2", '1')
	g.Connect("2", "3", '0')
	g.Connect("2", "3", '1')

	if ! g.Match("1110101010101001011000") {
		t.Fatalf("Did not match")
	}
}

func TestEmptyEdge(t *testing.T) {
	g := NewGraph([]string{"0", "1", "2", "3", "4"}, "0", []string{"1"})
	g.ConnectEmpty("0", "1")
	g.ConnectEmpty("0", "2")
	g.ConnectEmpty("2", "3")
	g.ConnectEmpty("3", "2")
	


	empty := g.nodes["0"].EmptyReachable()

	if empty.Size() != 4 {
		t.Fatalf("Error in EmptyEdge %d", empty.Size())
	}

	
}

func TestMatchEmpty(t *testing.T) {
	g := NewGraph([]string{"q", "r", "s", "t"}, "q", []string{"t"})
	g.Connect("q", "r", '0')
	g.Connect("r", "s", '0')
	g.Connect("s", "t", '0')
	g.Connect("t", "t", '0')
	g.Connect("q", "r", '1')
	g.Connect("r", "s", '1')
	g.Connect("r", "s", '1')
	g.Connect("s", "t", '1')
	g.Connect("t", "t", '1')
	g.ConnectEmpty("q", "r")
	g.ConnectEmpty("s", "t")

	if !g.Match("11") {
		t.Fatalf("No Match")
	}
}

func TestMatchLoop(t *testing.T) {
	g := NewGraph([]string{"q1", "q2", "q3"}, "q1", []string{"q3"})
	g.ConnectEmpty("q1", "q2")
	g.ConnectEmpty("q2", "q3")
	g.ConnectEmpty("q3", "q1")

	if !g.Match("") {
		t.Fatalf("No Match")
	}

}