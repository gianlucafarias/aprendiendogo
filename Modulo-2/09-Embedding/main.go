package main

import "fmt"

type Empleado struct {
	nombre string
	edad   int
}

type Gerente struct {
	Empleado
	departamento string
}

func (e Empleado) mostrarDetalles() {
	fmt.Println(e.nombre)
	fmt.Println(e.edad)
}

func (g Gerente) mostrarDetallesGerente() {
	fmt.Println(g.nombre)
	fmt.Println(g.edad)
	fmt.Println(g.departamento)
}

func main() {
	empleado1 := Empleado{
		nombre: "David",
		edad:   40,
	}
	empleado1.mostrarDetalles()

	gerente1 := Gerente{
		Empleado: Empleado{
			nombre: "Roman",
			edad:   60,
		},
		departamento: "Ventas",
	}

	gerente1.mostrarDetallesGerente()

}
