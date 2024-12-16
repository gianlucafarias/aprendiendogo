package main

import "fmt"

func main() {
	// Funcion que no usa closure

	counter := func() int {
		count := 0
		count++
		return count
	}
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// Funcion con closure
	counterClosure := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()
	fmt.Println(counterClosure())
	fmt.Println(counterClosure())
	fmt.Println(counterClosure())

}
