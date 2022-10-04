// Un Array es un contenedor de elementos de un tipo particular.
// Por ejemplo, y la matriz de enteros representa un contenedor de enteros.

package main

import "fmt"

func main() {
	var i [3]int
	i[0] = 42
	i[1] = 23
	i[2] = 17

	fmt.Println("0th index ", i[0])
	fmt.Println("1st index ", i[1])
	fmt.Println("2nd index ", i[2])
}
