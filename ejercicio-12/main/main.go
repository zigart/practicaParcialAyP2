package main

import (
	"fmt"
	"practica/funciones"
)

func main() {
	queue := funciones.NewQueue[int](0)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	resultado := funciones.Invertir(queue)
	fmt.Print(resultado)
}
