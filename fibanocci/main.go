package main

import "fmt"

func main(){

	a := 0
	b := 1
	var c int
	var sum int
	for  i := 2;i<10;i++ {
		c = a + b
		if i % 2 != 0 {
			sum += c
		}
		a = b
		b = c
	}

	fmt.Println(sum)

}
