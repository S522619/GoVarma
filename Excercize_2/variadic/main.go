package main

import "fmt"

func great(numbers ...int) int{
	var a int
	for _, n := range numbers{
	if n > a{
		a = n
	}
	}
	return a
}

func main(){
	number := great( 1,2,25,4,5,26,7,8)
	fmt.Println(number)
}
