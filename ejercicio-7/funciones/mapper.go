package funciones

import "practica/dictionary"

// func Mapper[T comparable](entrada dictionary.Dictionary[T, []T]) *dictionary.Dictionary[T, []T] {
// 	alumnos := dictionary.NewDictionary[string, []string]()

// 	for _, apellido := range entrada.GetKeys() {
// 		for _, alumno := range entrada.Get(apellido) {
// 			arreglo := []string{}
// 			if alumnos.Contains(alumno) {
// 				arreglo = alumnos.Get(alumno)
// 			}
// 			arreglo = append(arreglo, apellido)
// 			alumnos.Put(alumno, arreglo)
// 		}
// 	}
// 	return alumnos
// }

func Mapper(materias dictionary.Dictionary[string, []string]) dictionary.Dictionary[string, []string] {
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
