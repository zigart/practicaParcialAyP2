package main

import (
	"fmt"
	"practica/funciones"
	"practica/linkedlist"
)

func main() {
	list := linkedlist.NewLinkedList[int]()
	list.InsertAt(1, list.Size())
	list.InsertAt(3, list.Size())
	list.InsertAt(5, list.Size())

	list2 := linkedlist.NewLinkedList[int]()
	list2.InsertAt(2, list2.Size())
	list2.InsertAt(4, list2.Size())
	result := funciones.Intercalar(list, list2)
	fmt.Printf("%v", result)
}
