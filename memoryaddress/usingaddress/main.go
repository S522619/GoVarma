package main

import "fmt"

const kgtopound = 2

func main() {
	var kg int
	fmt.Println(" Enter the number of kgs ")
	fmt.Scan(&kg)
	pound := kg * kgtopound
	fmt.Printf("The value in pound's is : "+"%d", pound)
}
