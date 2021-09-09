package krus

// type nodeSet = map[*node]bool
type nodeMap = map[rune]*node

type node struct {
	name  string
	edges nodeMap
}

func newNode(name string) *node {
	return &node{name, make(nodeMap)}
}

// Finite state machine

type StateMachine struct {
	nodes  map[string]*node
	start  *node
	accept []*node
}

func NewGraph(nodeNames []string, start string, acceptStates []string) StateMachine {
	graph := StateMachine{make(map[string]*node), nil, []*node{}}

	for _, name := range nodeNames {
		graph.nodes[name] = newNode(name)
	}

	graph.start = graph.nodes[start]

	for _, name := range acceptStates {
		graph.accept = append(graph.accept, graph.nodes[name])
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

	// Check if accept state
	for _, node := range graph.accept {
		if node == currentNode {
			return true
		}
	}
	return false
}
