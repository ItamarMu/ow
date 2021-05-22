package tree

import (
	"ow_test/entities"
	"sync"
)

type Tree struct {
	children    []*node
	nodeMap     map[string]*node
	childrenMap map[string][]*node
	mapLock     sync.Mutex
}

func NewTree() *Tree {
	return &Tree{
		children:    make([]*node, 0),
		nodeMap:     make(map[string]*node),
		childrenMap: make(map[string][]*node),
	}
}

func (t *Tree) Add(msg *entities.Msg) {
	addedNode := newNode(msg)
	t.mapLock.Lock()
	defer t.mapLock.Unlock()
	t.nodeMap[addedNode.ID] = addedNode
	addedNode.children = t.childrenMap[addedNode.ID]
	if addedNode.ParentID == `` {
		t.children = append(t.children, addedNode)
	}
	t.childrenMap[addedNode.ParentID] = append(t.childrenMap[addedNode.ParentID], addedNode)
	_, parentExists := t.nodeMap[addedNode.ParentID]
	if parentExists {
		t.nodeMap[addedNode.ParentID].children = t.childrenMap[addedNode.ParentID]
	}
}

// An alternative Add implementation
func (t *Tree) Add2(msg *entities.Msg) {
	t.mapLock.Lock()
	defer t.mapLock.Unlock()
	addedNode := newNode(msg)
	_, exists := t.nodeMap[addedNode.ID]
	_, parentExists := t.nodeMap[addedNode.ParentID]
	if !parentExists {
		t.nodeMap[addedNode.ParentID] = &node{addedNode.ParentID, "", nil, make([]*node, 0)}
	}
	if exists {
		t.nodeMap[addedNode.ID].data = addedNode.data
		t.nodeMap[addedNode.ID].ParentID = addedNode.ParentID
	} else {
		t.nodeMap[addedNode.ID] = addedNode
	}
	if addedNode.ParentID == `` {
		t.children = append(t.children, t.nodeMap[addedNode.ID])
	} else { // not a first level node
		t.nodeMap[addedNode.ParentID].children = append(t.nodeMap[addedNode.ParentID].children, t.nodeMap[addedNode.ID])
	}
}
