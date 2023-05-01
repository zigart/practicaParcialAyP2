package funciones

type Registro struct {
	meses [12]MapaDeBits
}

func NewRegistro() *Registro {
	return &Registro{}
}

func stringToInt(mes string) int {
	switch mes {

	case "enero":
		return 0
	case "febrero":
		return 1
	case "marzo":
		return 2
	case "abril":
		return 3
	case "mayo":
		return 4
	case "junio":
		return 5
	case "julio":
		return 6
	case "agosto":
		return 7
	case "septiembre":
		return 8
	case "octubre":
		return 9
	case "noviembre":
		return 10
	case "diciembre":
		return 11
	}

	return -1
}

func (r *Registro) Registrar(mes string, dias ...uint8) {
	i := stringToInt(mes)
	if i < 0 {
		return
	}

	for _, dia := range dias {
		r.meses[i].Encender(dia)
	}
}

func (r *Registro) LLuvias(mes string) []uint8 {
	mesToInt := stringToInt(mes)
	arreglo := []uint8{}
	for i := uint8(0); i < 32; i++ {
		encendido, _ := r.meses[mesToInt].EstaEncendido(i)
		if encendido {
			arreglo = append(arreglo, i)
		}
	}
	return arreglo
}

func (r *Registro) LlovioAmbosMeses(primerMes, segundoMes string) []uint8 {
	lluviasPrimerMes := r.LLuvias(primerMes)
	lluviasSegundoMes := r.LLuvias(segundoMes)
	conjunto := NewSet[uint8]()

	for _, diasPrimerMes := range lluviasPrimerMes {
		for _, diasSegundoMes := range lluviasSegundoMes {
			if diasPrimerMes == diasSegundoMes {
				conjunto.Add(diasPrimerMes)
			}
		}
	}
	valoresCoincidencia := conjunto.Values()
	return valoresCoincidencia
}
