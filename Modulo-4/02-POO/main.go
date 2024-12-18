package main

import "fmt"

/*
- En GO no hay clases, tampoco herencia.
- Se puede reemplazar con estructuras y composicion para reutilizar codigo.
- En GO hay interfaces implicitas
- Polimorfismo Implicitio (no hace falta declarar explicitamente que un tipo cumple con una interfaz)
*/

type Hablador interface {
	Hablar()
}

type Persona struct {
	Nombre string
	Edad   int
}

func (p *Persona) Hablar() {
	fmt.Println("hola soy", p.Nombre)
}

func main() {
	p := Persona{"Juan", 30}
	p.Hablar()

}
