package main

import "fmt"

func main() {
	x := 10

	var puntero *int
	puntero = &x
	fmt.Println("La variable x es:", x)
	fmt.Println("La Direccion en memoria de la variable x es:", puntero)
	fmt.Println("El valor de x a traves del puntero es:", *puntero)

	modificarValor(&x)
	fmt.Println("El valor de X ahora es:", x)
}

// Funciones
func modificarValor(puntero *int) {
	fmt.Println("Valor de x modificado por puntero")
	*puntero = 50
}
