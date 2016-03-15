package main

import "fmt"

func main(){
	age := 44
	fmt.Println(age)
	fmt.Println(&age)

	abc(&age)

	fmt.Println(&age)
	fmt.Println(age)

}

func abc(z *int){
	fmt.Println(z)
	fmt.Println(*z)

	*z = 12

	fmt.Println(z)
	fmt.Println(*z)

}