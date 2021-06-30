package main

import "fmt"
import "time"

func main(){
	for i := 0; i<= 10;i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}