package funciones

import (
	"practica/linkedlist"
	"practica/stack"
)

func FindMax(list linkedlist.LinkedList[int]) int {
	stack := stack.NewStack[int](3)
	listValue := list.Get(0)
	stack.Push(listValue)
	for i := 0; i < list.Size(); i++ {
		val, _ := stack.Front()
		listValue = list.Get(i)
		if val < listValue {
			stack.Pop()
			stack.Push(listValue)
		}
	}
	max, _ := stack.Front()
	return max
}
