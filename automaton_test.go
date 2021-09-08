package krus

import "testing"

func TestNewNode(t *testing.T) {
	n := NewNode("q1")
	result := n.name
	expected := "q1"
	if result != expected {
		t.Fatalf(`NewNode("q1").name = %s, exptected %s`, result, expected)
	}
}