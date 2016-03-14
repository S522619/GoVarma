package main

import "fmt"

func abc() func() string{
	return func() string{
		return "Sandeep"
	}
}

func main(){
	a := abc()
	fmt.Println(a())
}
