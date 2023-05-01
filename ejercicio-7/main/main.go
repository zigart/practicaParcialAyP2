package main

import (
	"fmt"
	"practica/dictionary"
	"practica/funciones"
)

func main() {
	d1 := dictionary.NewDictionary[string, []string]()
	l1 := []string{"a", "b", "c"}
	l2 := []string{"d", "e", "f"}
	d1.Put("pepe", l1)
	d1.Put("jose", l2)

	fmt.Print(funciones.Mapper(d1))

}

/*
func Materias(materias dictionary.Dictionary[string, []string]) dictionary.Dictionary[string, []string] {
	alumnos := dictionary.NewDictionary[string, []string]()

	for _, apellido := range materias.GetKeys() {
		for _, alumno := range materias.Get(apellido) {
			arreglo := []string{}
			if alumnos.Contains(alumno) {
				arreglo = alumnos.Get(alumno)
			}
			arreglo = append(arreglo, apellido)
			alumnos.Put(alumno, arreglo)
		}
	}
	return alumnos
}
*/
