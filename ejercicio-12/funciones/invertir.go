package funciones

func Invertir[T comparable](queue *Queue[T]) *Stack[T] {

	if queue.IsEmpty() {
		return NewStack[T](0)
	}

	poppedValue, _ := queue.Dequeue()

	stackDeRecursion := Invertir(queue)

	stackDeRecursion.Push(poppedValue)
	return stackDeRecursion
}
