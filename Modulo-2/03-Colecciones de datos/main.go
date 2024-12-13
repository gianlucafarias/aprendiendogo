package main

import "fmt"

func main() {
	// Arreglos
	var nombres [3]string
	nombres[0] = "Gian"
	nombres[1] = "Guille"
	nombres[2] = "Ale"

	apellidos := [3]string{"Farias", "Bianchi", "Gonella"}
	fmt.Println(apellidos)

	// Slices
	var frutas []string
	frutas = append(frutas, "Manzana", "Pera", "Kiwi", "Naranja")

	verduras := []string{"Lechuga", "Brocoli", "Cebolla"}

	fmt.Println(frutas)
	fmt.Println(verduras)
}
