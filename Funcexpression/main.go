package main

import "fmt"

func main(){

	a := func(){
		fmt.Println("hello world!!")
	}
	a()
	fmt.Printf("%T",a)

}
