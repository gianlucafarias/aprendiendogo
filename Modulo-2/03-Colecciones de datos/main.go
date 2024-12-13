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

	// Maps
	var edades map[string]int
	edades = make(map[string]int)
	edades["Gian"] = 27
	edades["Guille"] = 28
	edades["Ale"] = 31

	edades2 := map[string]int{
		"Gian":   27,
		"Guille": 28,
		"Ale":    31,
	}

	fmt.Println(edades["Gian"])
	fmt.Println(edades2["Ale"])

	// Range
	numerosRange := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

	for i, num := range numerosRange {
		fmt.Println(i, num)
	}

	// Range map
	capitales := map[string]string{
		"Santa Fe":   "Santa Fe",
		"Entre Rios": "Parana",
		"Chaco":      "Resistencia",
	}

	for pais, capital := range capitales {
		fmt.Println(pais, capital)
	}
}
