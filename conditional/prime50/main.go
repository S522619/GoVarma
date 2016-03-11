package main

import "fmt"

func main(){
	for i:=0;i<50;i++ {
		if (i % 2 == 0 && i != 2) ||(i % 7 == 0 && i != 7)||(i % 3 == 0 && i != 3)||(i % 5 == 0 && i != 5){
			continue
		}


			fmt.Printf("%d \t", i)
	}
}
