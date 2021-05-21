package tree

import (
	"ow_test/entities"
)

var g int

type Tree struct {
	children []*node
	nodeMap  map[string]*node
}

func NewTree() *Tree {
	return &Tree{
		children: make([]*node, 0),
		nodeMap:  make(map[string]*node),
	}
}

func (t *Tree) Add(msg *entities.Msg) {
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
	} else { //not main leaf
		t.nodeMap[addedNode.ParentID].children = append(t.nodeMap[addedNode.ParentID].children, t.nodeMap[addedNode.ID])
	}
}
