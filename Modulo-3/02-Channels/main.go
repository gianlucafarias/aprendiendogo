package main

import "fmt"

func main() {
	// canal sin buffer
	canal := make(chan string)

	// enviar datos al canal
	// canal <- dato

	// recibir datos del canal
	// dato := <- canal

	// Goroutine para enviar un mensaje
	go func() {
		canal <- "Hola mundo desde Goroutine"
	}()

	// Recibiendo el mensaje
	mensaje := <-canal

	fmt.Println(mensaje)

	// canal con buffer
	canal2 := make(chan string, 2)

	canal2 <- "Hola mundo desde un canal"
	canal2 <- "Hola mundo desde un canal 2"

	fmt.Println(<-canal2)
	fmt.Println(<-canal2)

	// cerrar el canal
	close(canal2)

}
