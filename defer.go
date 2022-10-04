// En Go podemos usar un concepto llamado Aplazamiento para diferir la llamada de la función hasta que la función de cierre se haya procesado.
package main

import "fmt"

func timeMachine(i int) {
	fmt.Println("You have entered warp zone with number ", i)
}

func main() {
	i := 10000
	defer timeMachine(i)

	i = 20
	fmt.Println("Main ends here")
}
