package main

import "fmt"

func main(){
	var abc = map[int]string{
	1 : "Sandeep",
	2 : "Dilip",
		3:"Baby",
		4:"Daddy",
	}
fmt.Println(abc)
	if val,exists := abc[2]; exists{
		delete(abc,2)
		fmt.Println(val)

	} else {
		fmt.Println("The value doesnot exist !!")
	}


	fmt.Println(abc)







}
