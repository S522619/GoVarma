package main

import "fmt"

const (
	a  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {
	fmt.Println("Binary \t \t \t Decimal")
	fmt.Printf("%b \t \t \t ", KB)
	fmt.Printf("%d \t \t \t \n", KB)
	fmt.Printf("%b \t \t \t ", MB)
	fmt.Printf("%d \t \t \t \n", MB)
	fmt.Printf("%b \t \t ", GB)
	fmt.Printf("%d \t \t \t ", GB)

}
