// Operadores de postfix: los operadores de incremento y decremento se utilizan para agregar el valor de uno o eliminar el valor de uno, respectivamente, como se muestra a continuación.

package main

import "fmt"

func main() {
	var i int = 1
	i++
	fmt.Println("i++ = ", i)

	var j int = 1
	j--
	fmt.Println("j-- = ", j)
}
