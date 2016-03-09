package main

import "fmt"


func main(){

	//var a string = "I am a very good boy"
	x:= 1
	increment := func () int{
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())




}



