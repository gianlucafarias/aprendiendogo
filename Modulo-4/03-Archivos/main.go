package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// nombre del archivo que vamos a leer
	filename := "test.txt"

	// leer el archivo
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("El archivo no se puede leer", err)
		return
	}

	fmt.Println("Contenido del archivo:")
	fmt.Println(string(content))

	// ESCRIBIR ARCHIVO
	filenameNew := "test2.txt"
	text := "Hola este es un texto creado en un archivo desde Go"
	file, err := os.Create(filenameNew)
	if err != nil {
		fmt.Println("Ocurrió un error al crear el archivo", err)
		return
	}
	fmt.Println("Archivo creado con exito!")

	defer file.Close()

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println("Ocurrió un error al escribir el archivo")
	}
	fmt.Println("Archivo modificado con exito!")

	contentNew, err := ioutil.ReadFile(filenameNew)
	if err != nil {
		fmt.Println("El archivo no se puede leer", err)
		return
	}

	fmt.Println("Contenido del archivo Creado:")
	fmt.Println(string(contentNew))
}
