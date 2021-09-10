package krus

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

func (set nodeSet) Size() int {
	return len(set.storage)
}

func (set nodeSet) String() string {
	result := "["
	for n, _ := range set.storage {
		result += n.name + ", "
	}
	return result[:len(result)-2] + "]"
}


func (set nodeSet) ContainsAcceptNode() bool {
	for n, contained := range set.storage {
		if contained && n.isAccept {
			return true
		}
	}
	return false
}

func (set *nodeSet) InsertReachable() {
	newNodes := newNodeSet()
	for n, _ := range set.storage {
		newNodes.InsertSet(n.EmptyReachable())
	}
	set.InsertSet(newNodes)
}