package main

import "fmt"

func main() {
	/*var i int
	for i = 65; i < 123;i++ {
		//fmt.Println(i,"  ",string(i),"   ",[]byte(string(i)))
		fmt.Printf("%d  -  %v  -  %v \n",i,string(i),[]byte(string(i)))
	}*/
	foo := 'a'
	fmt.Printf("%T \n",foo)
	fmt.Printf("%d",rune(foo))
}
