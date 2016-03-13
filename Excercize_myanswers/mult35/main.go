package main

import "fmt"

func main() {

	var k = 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			k = k + i
		}
	}
	fmt.Printf("The total value is %d", k)
}
