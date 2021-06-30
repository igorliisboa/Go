package main

import "fmt"

func main(){

	if a := 10; a < 5{
		fmt.Println("a é menor que 5.")
	} else if a == 5{
		fmt.Println("a é igual a 5.")
	} else {
		fmt.Println("a é maior que 5.")
	}
}