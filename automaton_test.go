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
	if q0.edges['1'] != q1 {
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
