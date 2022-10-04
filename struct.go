// Struct en Golang es una recopilación de datos.
// Podemos usarlo para definir tipos de datos personalizados.
// Una estructura puede contener campos miembros dentro de ellos e incluso puede contener otros tipos de estructura.
// El siguiente código muestra una estructura de ejemplo en Golang para un planeta.
package main

import "fmt"

type Planetas struct {
	nbPlaneta string
	nbGalaxia string
	nuLunas   int
}

func main() {
	tierra := Planetas{"Tierra", "Via Lactea", 1}
	fmt.Println(tierra)

	jupiter := Planetas{"Jupiter", "Via Lactea", 21}
	fmt.Println(jupiter)

	//Podemos acceder y cambiar los campos de estructura utilizando el operador de punto (.). También podemos dar el nombre de inicialización de los campos para una mejor legibilidad.

	tierra.nbGalaxia = "Andromeda"
	fmt.Println(tierra)
}
