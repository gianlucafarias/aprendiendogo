package main

import (
	"fmt"
)

// para que reconozca como paquete a "matematicas" se tiene que hacer go mod init mi_proyecto en el root del proyecto

func main() {
	suma := matematicas.suma(3 + 2)
	resta := matematicas.resta(3 - 2)

	fmt.Println(suma)
	fmt.Println(resta)
}
