package main

import "fmt"

func main(){
	m := make([]string,4,22)  // length = 4 and capacity = 22
	// fmt.Println(m)
	fmt.Println(cme(m))

}

func cme(x [] string) string{
	x[0] = "Sandeep"
	return x[0]
}
