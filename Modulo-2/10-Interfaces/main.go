package main

import "fmt"

// Declarar interface
type Figura interface {
	Area() float64
	Perimetro() float64
}

type Circulo struct {
	radio float64
}

func (c Circulo) Area() float64 {
	return 3.1416 * c.radio * c.radio
}

func main() {

	circulo1 := Circulo{
		radio: 250,
	}

	fmt.Println("El area del circulo es:", circulo1.Area())
}
