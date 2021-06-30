package main

import "fmt"

func main(){
	for a := 0;a <= 10;a ++{
		fmt.Println(a)
	}

	var soma int = 2
	for {
		soma += soma
		fmt.Println(soma)
		if soma == 1024{
			break
		}
	}

}