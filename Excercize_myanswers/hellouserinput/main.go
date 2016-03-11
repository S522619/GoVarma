package main

import "fmt"

func main(){
	var i string
	fmt.Print("Enter the name:  ")
	fmt.Scan(&i)
	fmt.Printf("I have entered %s",i)
}