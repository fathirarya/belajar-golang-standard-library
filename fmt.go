package main

import "fmt"

func main() {
	firstName := "M Fathir"
	LastName := "Arya Nafis"

	fmt.Println("Hello '", firstName, LastName, "'")
	
	fmt.Printf("Hello '%s %s'\n", firstName, LastName)
}