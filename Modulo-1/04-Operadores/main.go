package main

import "fmt"

func main() {
	// Operadores Arimeticos

	num1 := 10
	num2 := 5

	suma := num1 + num2
	resta := num1 - num2
	multiplicacion := num1 * num2
	division := num1 / num2
	modulo := num1 % num2

	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
	fmt.Println("Multiplicacion:", multiplicacion)
	fmt.Println("Division:", division)
	fmt.Println("Modulo:", modulo)

	// Operadores Logicos
	/*  AND &&
	OR ||
	NOT !
	*/

	esAdulto := true
	tienePermiso := false

	puedeEntrar := esAdulto && tienePermiso
	puedeSalir := esAdulto || tienePermiso
	noPuedeSalir := !tienePermiso

	fmt.Println("Puede entrar:", puedeEntrar)
	fmt.Println("Puede salir:", puedeSalir)
	fmt.Println("No puede salir:", noPuedeSalir)

}
