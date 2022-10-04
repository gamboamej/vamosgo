package main

import "fmt"

func main() {
	// Convertir entero en Decimal: En el siguiente ejemplo, estamos convirtiendo Integer a Float.
	var int_numero int = 8
	fmt.Printf("Valor = %v, Tipo = %T\n", int_numero, int_numero)

	var flo_numero float32 = float32(int_numero) + 1.8
	fmt.Printf("Valor = %v, Tipo = %T\n", flo_numero, flo_numero)
	// Convertir float a entero: En el siguiente ejemplo, estamos convirtiendo de Float a Integer.
	int_numero = int(flo_numero)
	fmt.Printf("Valor = %v, Tipo = %T\n", int_numero, int_numero)
}
