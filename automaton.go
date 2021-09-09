package krus

type nodeMap = map[rune]*node

type node struct {
	name  string
	edges nodeMap
	isAccept bool
}

func newNode(name string) *node {
	return &node{name, make(nodeMap), false}
}

type nodeSet struct {
	storage map[*node]bool
}

func newNodeSet() nodeSet {
	return nodeSet{ make(map[*node]bool) }
}

func (set *nodeSet) Insert(newNode *node) {
	set.storage[newNode] = true
}

func (set *nodeSet) InsertSet(other nodeSet) {
	for n, _ := range other.storage {
		set.storage[n] = true
	}
}

func (set *nodeSet) Contains(n *node) bool {
	return set.storage[n]
}

// Finite state machine

type StateMachine struct {
	nodes  map[string]*node
	start  *node
}

func NewGraph(nodeNames []string, start string, acceptStates []string) StateMachine {
	graph := StateMachine{make(map[string]*node), nil }

	for _, name := range nodeNames {
		graph.nodes[name] = newNode(name)
	}

	graph.start = graph.nodes[start]

	for _, name := range acceptStates {
		graph.nodes[name].isAccept = true
	}

	return graph
}

func (graph *StateMachine) Connect(source string, target string, symbol rune) {
	sourceNode := graph.nodes[source]
	targetNode := graph.nodes[target]
	sourceNode.edges[symbol] = targetNode
}

func (graph StateMachine) Match(tomatch string) bool {
	currentNode := graph.start
	for _, symbol := range tomatch {
		currentNode = currentNode.edges[symbol]
	}

	return currentNode.isAccept
}
