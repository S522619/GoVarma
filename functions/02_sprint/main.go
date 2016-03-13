package main

import "fmt"

func main(){
fmt.Println(abc("Sandeep ","Dilip "))
}


func abc(a,b string) (s string, f string){
 	s = fmt.Sprint(a, b);f =fmt.Sprint(b, a)
	return s,f
	}
