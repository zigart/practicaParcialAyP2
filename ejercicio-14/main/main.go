package main

import (
	"fmt"
	"practica/funciones"
)

func main() {
	arr1 := []int{1, 5, 3, 4}
	resultadoMax, resultadoMin := funciones.FindMaxAndMin(arr1)
	fmt.Print(resultadoMax)
	fmt.Print(resultadoMin)
}
