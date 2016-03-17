package main

import "fmt"

func Sqr(n int) int{

	var f,g int
	for i:=1;i<=n;i++{
		f = f + (i*i)
		g = g + i
	}
	return  g*g-f
}

func main(){
	x := Sqr(100)
	fmt.Println(x)
}
