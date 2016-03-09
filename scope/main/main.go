package main

import "fmt"

/*var x string = " Hi guys"
func increment() string {

	return x +" " + "abc"
}*/

func main(){
	x:= 1
	increment := func () int{
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())




}



