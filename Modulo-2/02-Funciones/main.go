package main

import "fmt"

func main() {
	numero1 := 10
	numero2 := 14

	resultado := sumar(numero1, numero2)

	fmt.Println(resultado)

	suma, resta := sumar_restar(numero1, numero2)
	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)

	saludar()
}

// Declaracion de una funcion
func sumar(num1 int, num2 int) int {
	return num1 + num2
}

// Funcion Multiple
func sumar_restar(num1 int, num2 int) (int, int) {
	suma := num1 + num2
	resta := num1 - num2
	return suma, resta
}

// Void
func saludar() {
	fmt.Println("Hola mundo!")
}
