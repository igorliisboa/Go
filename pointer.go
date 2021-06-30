package main

import "fmt"

func main(){
	a := 32
	var p *int = &a
	*p = *p + 2
	fmt.Println(a)
}