package main

import "fmt"

// Definir Estructura
type Libro struct {
	titulo  string
	autor   string
	paginas int
	precio  float64
}

// Crear METODOS
// Metodo asociado a la estructura para mostrar detalles (por valor)
func (b Libro) mostrarDetalles() {
	fmt.Println("Titulo del Libro:", b.titulo)
	fmt.Println("Autor:", b.autor)
	fmt.Println("N* de Paginas:", b.paginas)
	fmt.Println("Precio: ", b.precio)
}

// Metodo asociado a la estructura para modificar el objeto (por puntero)
func (b *Libro) aplicarDescuento(descuento float64) {
	b.precio = b.precio - descuento
}

func main() {

	miLibro := Libro{
		titulo:  "Programacion en go",
		autor:   "Gianluca Palmier",
		paginas: 70,
		precio:  2000.00,
	}

	// Receptor por valor
	miLibro.mostrarDetalles()

	// Receptor por puntero
	miLibro.aplicarDescuento(50)
	fmt.Println("Descuento Aplicado:", miLibro.precio)

}
