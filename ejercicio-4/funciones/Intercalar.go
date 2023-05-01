package funciones

import (
	"practica/linkedlist"
)

func Intercalar[T comparable](l1, l2 *linkedlist.LinkedList[T]) *linkedlist.LinkedList[T] {

	if l1.Size() == 0 && l2.Size() != 0 {
		return l2
	} else if l1.Size() != 0 && l2.Size() == 0 {
		return l1
	}

	newList := linkedlist.NewLinkedList[T]()

	for i := 0; i < l1.Size(); i++ {
		liValue := l1.Get(i)
		newList.InsertAt(liValue, newList.Size())
		if i < l2.Size() && !l2.IsEmpty() {
			newList.InsertAt(l2.Get(i), newList.Size())
		}
	}

	return newList
}
