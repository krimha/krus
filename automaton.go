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

func (this *node) Edges(symbol rune) *nodeSet {
	return this.edges[symbol]
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

func (graph *StateMachine) Node(nodeName string) *node {
	return graph.nodes[nodeName]
}

func (graph *StateMachine) Connect(source string, target string, symbol rune) {
	sourceNode := graph.Node(source)
	targetNode := graph.Node(target)
	if sourceNode.Edges(symbol) == nil {
		newSet := newNodeSet()
		sourceNode.edges[symbol] = &newSet
	}
	sourceNode.Edges(symbol).Insert(targetNode)
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
	currentNodeSet.InsertReachable()

	for _, symbol := range tomatch {
		newCurrentNodeSet := newNodeSet()
		for sourceNode, _ := range currentNodeSet.storage {
			// TODO Should get empty set instead?
			toInsert := sourceNode.Edges(symbol)
			if toInsert != nil {
				newCurrentNodeSet.InsertSet(*sourceNode.Edges(symbol))
			}
		}
		currentNodeSet = newCurrentNodeSet
		currentNodeSet.InsertReachable()
	}

	return currentNodeSet.ContainsAcceptNode()
}