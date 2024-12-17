package main

import (
	"fmt"
)

type DivisionError struct {
	Dividendo float64
	Divisor   float64
}

func (e *DivisionError) Error() string {
	return fmt.Sprintf("No se puede dividir por cero", e.Dividendo, e.Divisor)
}

func main() {
	// usar defer para garantizar el manejo del panic/recover
	defer manejarPanic()
	resultado, err := division(10.00, 0.0)
	if err != nil {
		fmt.Println("Error", err)
		return

	}
	fmt.Println(resultado)

}

func division(a float64, b float64) (float64, error) {
	if b == 0 {
		/*
			Con Errores New y mensaje personalizado

			return 0, errors.New("No se puede dividir por 0!")
		*/
		/*
			Con tipo de mensaje

			 return 0, &DivisionError{a, b}
		*/
		/*
			con PANIC (Detiene la ejecucion y lanza el error)
		*/
		panic("PANIC!! No se puede dividir por 0!")

	}
	return a / b, nil
}

func manejarPanic() {
	if r := recover(); r != nil {
		fmt.Println("Ocurrió un error:", r)
	}
}
