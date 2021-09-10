package krus

type nodeMap = map[rune]*nodeSet

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

func (set nodeSet) String() string {
	result := "["
	for n, _ := range set.storage {
		result += n.name + ", "
	}
	return result[:len(result)-2] + "]"
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
	if sourceNode.edges[symbol] == nil {
		newSet := newNodeSet()
		sourceNode.edges[symbol] = &newSet
	}
	sourceNode.edges[symbol].Insert(targetNode)
}

func (graph StateMachine) Match(tomatch string) bool {
	currentNodeSet := newNodeSet()
	currentNodeSet.Insert(graph.start)

	// TODO: Need to wrap this in another loop for the multiple nodes case
	for _, symbol := range tomatch {
		newCurrentNodeSet := newNodeSet()
		for sourceNode, _ := range currentNodeSet.storage {
			toInsert := sourceNode.edges[symbol]
			if toInsert != nil {
				newCurrentNodeSet.InsertSet(*sourceNode.edges[symbol])
			}
		}
		currentNodeSet = newCurrentNodeSet
	}

	for n, _ := range currentNodeSet.storage {
		if n.isAccept {
			return true
		}
	}
	return false
}
