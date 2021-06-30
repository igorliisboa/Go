package main

import "fmt"

func main(){
	nome := "Igor"
	switch nome{
	case "João":
		fmt.Println("Não é o Igor")
	case "Igor":
		fmt.Println("É o Igor!")
	default:
		fmt.Println("Não conheço!")
	}
}