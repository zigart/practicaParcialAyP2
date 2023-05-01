package main

import (
	"fmt"
	"practica/funciones"
)

func main() {
	registro := funciones.NewRegistro()
	registro.Registrar("enero", 10, 11, 12)
	registro.Registrar("febrero", 10)
	enero := registro.LLuvias("enero")
	fmt.Print(enero)
	coincidencia := registro.LlovioAmbosMeses("enero", "febrero")
	fmt.Print(coincidencia)
}
