package main

import "fmt"

func main() {
	var capitalPaises map[string]string
	capitalPaises = make(map[string]string)
	capitalPaises["India"] = "New Delhi"
	capitalPaises["U.S.A"] = "Washington DC"
	capitalPaises["Japan"] = "Tokyo"

	fmt.Println(capitalPaises)
	fmt.Println(capitalPaises["India"])
}
