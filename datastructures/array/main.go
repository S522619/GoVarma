package main

import "fmt"

func main(){
	var s [] string

	for i:=65;i<91;i++{

		s = append(s,string(i))
	}
	fmt.Println(s)
	fmt.Println(s[25])
}
