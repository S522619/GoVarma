package main

import "fmt"

func main(){
	i:=0
	for{
		i++
		if i%2 == 0{
			continue
		}
		fmt.Printf("%d \n",i)
		if i>=50{
			break
		}
	}
}
