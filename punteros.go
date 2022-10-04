package main

import "fmt"

func main() {
	message := "Hola Pratapo"

	// Pointer to string
	var pMessage *string

	// pMessage points to addr of message
	pMessage = &message
	fmt.Println("Message = ", *pMessage)
	fmt.Println("Message Address = ", pMessage)

	// Change message using pointer de-referencing
	*pMessage = "Hello Universe"
	fmt.Println("Message = ", *pMessage)
	fmt.Println("Message Address = ", pMessage)
}
