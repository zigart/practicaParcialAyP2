package main

import "practica/funciones"

func main() {
	Solitario := funciones.NuevoSolitario()
	carta := funciones.NuevaCarta("ORO", "K")
	carta2 := funciones.NuevaCarta("ORO", "q")
	carta3 := funciones.NuevaCarta("ORO", "j")
	Solitario.ApilarEnStack1(carta)
	Solitario.ApilarEnStack1(carta2)
	Solitario.ApilarEnStack1(carta3)
	Solitario.ImprimirPilas()
}
