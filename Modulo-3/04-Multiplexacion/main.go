package main

import "fmt"

// escuchar varios canales sin bloquear el programa

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		for {
			canal1 <- "Mensaje1"
		}
	}()

	go func() {
		for {
			canal2 <- "Mensaje1"
		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case mensaje1 := <-canal1:
			fmt.Println(mensaje1)
		case mensaje2 := <-canal2:
			fmt.Println(mensaje2)

		}
	}
}
