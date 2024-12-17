package main

import (
	"fmt"
	"time"
)

/*
Permite trabajar con muchos canales al mismo tiempo. Es como un switch de canales
*/

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	// goroutine para enviar datos al canal 1 despues de un segundo
	go func() {
		time.Sleep(time.Second)
		canal1 <- "Mensaje desde el canal 1"
	}()

	// goroutine para enviar datos al canal 2 despues de dos segundos
	go func() {
		time.Sleep(time.Second * 2)
		canal2 <- "Mensaje desde el canal 2"
	}()

	// uso del select para esperar mensajes de ambos canales
	for i := 0; i < 2; i++ {
		select {
		case mensaje1 := <-canal1: // si el canal 1 tiene un mensaje listo
			println(mensaje1)
		case mensaje2 := <-canal2: // si el canal 2 tiene un mensaje listo
			fmt.Println(mensaje2)
		}
	}
}
