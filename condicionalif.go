package main

import "fmt"

func main() {
	var balanceCuenta int = 0
	if balanceCuenta < 1000 {
		fmt.Println("Cuenta por cancelar!")
	}

	if balanceCuenta := 0; balanceCuenta < 1000 {
		fmt.Println("Cuenta Cerrada!")
	}

	if balanceCuenta := 1001; balanceCuenta < 1000 {
		fmt.Println("Cuenta Cerrada!")
	} else {
		fmt.Println("Nos encanta tenerte en Banco Pratapo!")
	}

	if balanceCuenta := 1000001; balanceCuenta < 1000 {
		fmt.Println("Cuenta Cerrada!")
	} else if balanceCuenta > 1000000 {
		fmt.Println("Encuentre un paquete de crucero turístico por Europa en su buzón de correo.")
	} else {
		fmt.Println("Nos encanta tenerte en Banco Pratapo!")
	}

	if balanceCuenta := 501; balanceCuenta < 1000 {
		if balanceCuenta < 500 {
			fmt.Println("Cuenta Cerrada!")
		} else {
			fmt.Println("¡Será mejor que mantengas un saldo mínimo, amigo! Tienes 5 días de tiempo")
		}
	} else if balanceCuenta > 1000000 {
		fmt.Println("Encuentre un paquete de crucero turístico por Europa en su buzón de correo.")
	} else {
		fmt.Println("Nos encanta tenerte en Banco Pratapo!")
	}

	var alphabet string = "A"
	switch alphabet {
	case "A":
		fmt.Println("Apple")
	case "B":
		fmt.Println("Ball")
	case "C":
		fmt.Println("Cat")
	case "D":
		fmt.Println("Doll")
	default:
		fmt.Println("EZ!")
	}

}
