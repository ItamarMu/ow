package main

import (
	"fmt"
	"ow_test/entities"
)

func main() {
	fmt.Println("asd")
	var x1 = entities.Msg{
		ID: "IDasd",
		ParentID: "parent",
		Data: "data",
	}
	fmt.Println(x1.ID)
}