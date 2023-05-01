package funciones

/*
A: 2
B: 2
C: 0

log 2 (2) > 0 =>


*/
func FindMaxAndMin(numeros []int) (max, min int) {

	if len(numeros) == 1 {
		return numeros[0], numeros[0]
	}

	mitad := len(numeros) / 2
	primeraMitad := numeros[:mitad]
	segundaMitad := numeros[mitad:]

	maxA, minA := FindMaxAndMin(primeraMitad)
	maxB, minB := FindMaxAndMin(segundaMitad)

	maximosYMinimos := [2]int{}
	if maxA > maxB {
		maximosYMinimos[0] = maxA
	} else {
		maximosYMinimos[0] = maxB
	}

	if minA < minB {
		maximosYMinimos[1] = minA
	} else {
		maximosYMinimos[1] = minB
	}

	return maximosYMinimos[0], maximosYMinimos[1]
}
