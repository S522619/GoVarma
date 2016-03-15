package main

import "fmt"

func abc(numbers []int, callback func(int)){
	for _,n:= range numbers{
		callback(n)
	}
}

func main(){
	abc([] int{21,24,25,15},func(nk int){
		fmt.Println(nk)
	})
}

