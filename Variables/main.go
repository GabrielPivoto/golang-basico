package main

import "fmt"

var nome string //declaração

func main() {

	nome = "Jaum di Kamargu"
	fmt.Println(nome)

	idade := 35 //declaração e atribuição
	fmt.Printf("%v \n", idade)

	pi := 3.14
	flag := true
	stringEstranha := `Uma string 
	
	estranha
	`

	//Print formatado
	fmt.Printf("%v \n", pi)
	fmt.Printf("%v \n", flag)
	fmt.Printf("%v \n", stringEstranha)
	fmt.Printf("\n")

	//Tipod das variaveis
	fmt.Printf("%T \n", pi)
	fmt.Printf("%T \n", flag)
	fmt.Printf("%T \n", stringEstranha)

}
