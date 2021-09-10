package krus

type nodeMap = map[rune]*nodeSet

type node struct {
	name  string
	edges nodeMap
	emptyEdges *nodeSet
	isAccept bool
}

func newNode(name string) *node {
	newSet := newNodeSet()
	return &node{name, make(nodeMap), &newSet, false}
}

func (this *node) EmptyReachable() nodeSet {
	result := newNodeSet()
	result.Insert(this)

	for {
		newNodes := newNodeSet()
		for currNode, _ := range result.storage {
			newNodes.InsertSet(*currNode.emptyEdges)
		}
		
		oldLength := result.Size()
		result.InsertSet(newNodes)
		// No new nodes were found
		if result.Size() == oldLength {
			break
		}
	}

	return result
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

func (graph *StateMachine) ConnectEmpty(source string, target string) {
	sourceNode := graph.nodes[source]
	targetNode := graph.nodes[target]
	if sourceNode.emptyEdges == nil {
		newSet := newNodeSet()
		sourceNode.emptyEdges = &newSet
	}
	sourceNode.emptyEdges.Insert(targetNode)
}

func (graph StateMachine) Match(tomatch string) bool {
	currentNodeSet := newNodeSet()
	currentNodeSet.Insert(graph.start)

	for _, symbol := range tomatch {
		newCurrentNodeSet := newNodeSet()
		for sourceNode, _ := range currentNodeSet.storage {
			// TODO Should get empty set instead?
			toInsert := sourceNode.edges[symbol]
			if toInsert != nil {
				newCurrentNodeSet.InsertSet(*sourceNode.edges[symbol])
			}
		}
		currentNodeSet = newCurrentNodeSet
	}

	return currentNodeSet.ContainsAcceptNode()
}