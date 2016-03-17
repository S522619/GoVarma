package main

import "fmt"

/*
func xyz(n int) (float64, bool){
	return float64(n/2), n%2 == 0
}

func main(){
	x,y := xyz(25)
	fmt.Println(x,y)
}*/

func main(){
	xyz := func(n int)(int, bool){
		return n/2, n%2==0
	}
	fmt.Println(xyz(25))

}