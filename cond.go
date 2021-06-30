package main

import "fmt"

func main(){
	var a int = 3
	var b int = 4

	if a < b{
		fmt.Println("a é menor que b.")
	} else if a == b{
		fmt.Println("a é igual a b.")
	} else {
		fmt.Println("a é maior que b.")
	}
}