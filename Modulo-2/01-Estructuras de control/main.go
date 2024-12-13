package main

import "fmt"

func main() {
	// Estructuras de control en Go

	edad := 19
	// Condicionales | IF ELSE SWITCH

	if edad < 18 {
		fmt.Println("El usuario es menor de edad")
	} else {
		fmt.Println("El usuario es mayor de edad")
	}

	dia := 3
	switch dia {
	case 1:
		fmt.Println("Es Domingo")
	case 2:
		fmt.Println("Es Lunes")
	case 3:
		fmt.Println("Es Martes")
	case 4:
		fmt.Println("Es Miercoles")
	case 5:
		fmt.Println("Es Jueves")
	case 6:
		fmt.Println("Es Viernes")
	case 7:
		fmt.Println("Es Sabado")
	default:
		fmt.Println("Dia no valido")
	}

	// Bucles

	/*
		FOR clasico
	*/

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	/*
		FOR con rango
	*/

	nombres := []string{"Gian", "Guille", "Ale"}

	for i, nombre := range nombres {
		fmt.Println(i, nombre)
	}
}
