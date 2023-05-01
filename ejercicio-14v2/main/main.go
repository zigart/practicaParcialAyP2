package main

import (
	"fmt"
	"practica/funciones"
)

func main() {
	max, min := funciones.FindMaxAndMin([]int{1, 2, 10, 4, 5, 6, 7, 8})
	fmt.Printf("max %v \n", max)
	fmt.Printf("min %v \n", min)
}
