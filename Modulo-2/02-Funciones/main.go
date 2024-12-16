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

	// Funcion Avanzada
	multiplicar := func(a int, b int) int {
		return a * b
	}

	resultadoMultiplicacion := multiplicar(3, 4)

	fmt.Println("El resultado de la multiplicación es:", resultadoMultiplicacion)

	// Crear una funcion e imprimirla
	doble := func(a int) int {
		return a * 2
	}(5)
	fmt.Println(doble)

	// Usar una funcion anonima sin asignarla a una variable
	fmt.Println("La suma es:", func(x int, y int) int {
		return x + y
	}(10, 5))

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

// Funcion que retorna otra funcion
func multiplicador(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}
