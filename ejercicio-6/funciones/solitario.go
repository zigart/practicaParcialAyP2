package funciones

import (
	"fmt"
	"strings"
)

type Carta struct {
	palo  string
	valor string
}

type Solitario struct {
	pila1 Stack[Carta]
	pila2 Stack[Carta]
	pila3 Stack[Carta]
	pila4 Stack[Carta]
}

func NuevoSolitario() *Solitario {
	return &Solitario{}
}

func NuevaCarta(palo, valor string) *Carta {
	return &Carta{palo: palo, valor: valor}
}

func comprobarValor(cartaDePila, cartaAAgregar *Carta) bool {
	diccionario := NewDictionary[string, string]()
	diccionario.Put("K", "")
	diccionario.Put("Q", "K")
	diccionario.Put("J", "Q")
	diccionario.Put("10", "J")
	diccionario.Put("9", "10")
	diccionario.Put("8", "9")
	diccionario.Put("7", "8")
	diccionario.Put("6", "7")
	diccionario.Put("5", "6")
	diccionario.Put("4", "5")
	diccionario.Put("3", "4")
	diccionario.Put("2", "3")
	diccionario.Put("A", "2")
	cartaAAgregarPalo := strings.ToUpper(cartaAAgregar.palo)
	paloStack := strings.ToUpper(cartaDePila.palo)
	cartaAAgregarValor := strings.ToUpper(cartaAAgregar.valor)
	valorStack := strings.ToUpper(cartaDePila.valor)
	if paloStack != cartaAAgregarPalo {
		return false
	}
	valorDic := diccionario.Get(cartaAAgregarValor)
	valorDic = strings.ToUpper(valorDic)

	return valorDic == valorStack
}

// nuevoSolitario := &newSolitario[]()
// nuevoSolitario.ApilarEnStack1()
func (Solitario *Solitario) ApilarEnStack1(carta *Carta) {
	if Solitario.pila1.IsEmpty() {
		Solitario.pila1.Push(*carta)
	}
	valorDePila, _ := Solitario.pila1.Front()
	comprobarValor := comprobarValor(&valorDePila, carta)
	if comprobarValor {
		Solitario.pila1.Push(*carta)
	} else {
		fmt.Printf("no apilo \n")
	}
}
func (Solitario *Solitario) ApilarEnStack2(carta *Carta) {
	if Solitario.pila2.IsEmpty() {
		Solitario.pila2.Push(*carta)
	}
	valorDePila, _ := Solitario.pila2.Front()
	comprobarValor := comprobarValor(&valorDePila, carta)
	if comprobarValor {
		Solitario.pila2.Push(*carta)
	} else {
		fmt.Printf("no apilo \n")
	}
}
func (Solitario *Solitario) ApilarEnStack3(carta *Carta) {
	if Solitario.pila3.IsEmpty() {
		Solitario.pila3.Push(*carta)
	}
	valorDePila, _ := Solitario.pila3.Front()
	comprobarValor := comprobarValor(&valorDePila, carta)
	if comprobarValor {
		Solitario.pila3.Push(*carta)
	} else {
		fmt.Printf("no apilo \n")
	}
}
func (Solitario *Solitario) ApilarEnStack4(carta *Carta) {
	if Solitario.pila4.IsEmpty() {
		Solitario.pila4.Push(*carta)
	}
	valorDePila, _ := Solitario.pila4.Front()
	comprobarValor := comprobarValor(&valorDePila, carta)
	if comprobarValor {
		Solitario.pila4.Push(*carta)
	} else {
		fmt.Printf("no apilo")
	}
}

func (solitario *Solitario) ImprimirPilas() {
	front1, _ := solitario.pila1.Front()
	fmt.Printf("%v \n", front1)
}
