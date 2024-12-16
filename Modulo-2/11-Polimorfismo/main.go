package main

import "fmt"

type Forma interface {
	Area() float64
}

type Circulo struct {
	radio float64
}

type Rectangunlo struct {
	ancho float64
	alto  float64
}

func (r Rectangunlo) Area() float64 {
	return r.alto * r.ancho
}

func (c Circulo) Area() float64 {
	return 3.1516 * c.radio * c.radio
}

func main() {

	circulo1 := Circulo{
		radio: 20,
	}

	rectangulo1 := Rectangunlo{
		ancho: 200,
		alto:  300,
	}

	fmt.Println("El Area del Rectangulo es:", rectangulo1.Area())
	fmt.Println("El Area del Circulo es:", circulo1.Area())

}
