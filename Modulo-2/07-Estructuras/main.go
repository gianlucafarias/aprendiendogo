package main

import "fmt"

func main() {
	var persona1 Persona

	persona1.nombre = "Gianluca"
	persona1.edad = 27
	persona1.trabajo = "Municipalidad de Ceres"
	persona1.salario = 1000000

	persona2 := Persona{
		nombre:  "Guillermina",
		edad:    27,
		trabajo: "Municipalidad de Ceres",
		salario: 2000000,
	}

	fmt.Println(persona1)
	fmt.Println(persona2)
}

// Estructuras
type Persona struct {
	nombre  string
	edad    int
	trabajo string
	salario int
}
