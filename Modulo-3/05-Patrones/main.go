package main

import "fmt"

// combinar datos de multiples canales en un solo canal

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		canal1 <- "mensaje1"
		close(canal1)
	}()

	go func() {
		canal2 <- "mensaje2"
		close(canal2)
	}()

	canalCombinado := combinar(canal1, canal2)

	for mensaje := range canalCombinado {
		fmt.Println(mensaje)
	}

}

func combinar(canal1, canal2 <-chan string) <-chan string {
	salida := make(chan string)

	go func() {
		for mensaje := range canal1 {
			salida <- mensaje
		}
	}()

	go func() {
		for mensaje := range canal2 {
			salida <- mensaje
		}
	}()
	return salida
}
