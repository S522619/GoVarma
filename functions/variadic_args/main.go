package main

import "fmt"

func main(){
	data := []float64{21.2, 24.9, 26.5, 20.6,41.0}
	n := average(data...)
	fmt.Println(n)

}

func average(a ...float64)  float64{
	var total float64
for _,t := range a{
	total += t
}
	return total/float64(len(a))
}
