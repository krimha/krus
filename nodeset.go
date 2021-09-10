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

func (set nodeSet) String() string {
	result := "["
	for n, _ := range set.storage {
		result += n.name + ", "
	}
	return result[:len(result)-2] + "]"
}