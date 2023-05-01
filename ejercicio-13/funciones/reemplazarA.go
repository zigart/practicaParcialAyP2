package funciones

import (
	"fmt"
	"strings"
)

func HacerA(cadena string) string {
	cadena = strings.ToLower(cadena)

	if len(cadena) == 0 {
		return ""
	}

	ultimaLetra := string(cadena[len(cadena)-1])
	fmt.Printf("%v \n", ultimaLetra)
	if ultimaLetra == "e" || ultimaLetra == "i" || ultimaLetra == "o" || ultimaLetra == "u" {
		ultimaLetra = "a"
	}

	resto := cadena[:len(cadena)-1]
	HacerA(resto)
	fmt.Printf("%s", ultimaLetra)
	return ultimaLetra
}
