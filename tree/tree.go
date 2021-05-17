package tree

import (
	"fmt"
	"ow_test/entities"
)

var g int

type Tree struct {
	children []*node
}

func NewTree() *Tree {
	return &Tree{
		children: make([]*node, 0),
	}
}

func (t *Tree) Add(msg *entities.Msg) {
	if msg.ParentID == `` {
		t.children = append(t.children, newNode(msg))
	} else {
		for _, v := range t.children {
			AddRec(v, msg)
		}
	}
	g++
	if g%10000 == 0 {
		fmt.Printf("added = %d\n", g)
	}
}

func AddRec(node *node, msg *entities.Msg) {
	// fmt.Println("hi")
	if node.ID == msg.ParentID {
		node.children = append(node.children, newNode(msg))
	} else {
		for _, v := range node.children {
			AddRec(v, msg)
		}
	}

}

func GetParent(node *node) {

}
