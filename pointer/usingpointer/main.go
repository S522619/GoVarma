package main

import "fmt"

func zero(y *int){
	fmt.Println(y)
	*y = 8
}

func main(){
	x := 5
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)

}
