package main

import (
	"fmt"
	"practica/funciones"
	"practica/linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList[int]()
	list.InsertAt(1, list.Size())
	list.InsertAt(2, list.Size())
	list.InsertAt(3, list.Size())
	result := funciones.FindMax(*list)
	fmt.Printf("%v", result)
}
