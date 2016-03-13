package main

import "fmt"

func main(){
fmt.Println(average(5, 7, 89))
}

func average(a ...float64) float64{
	/*fmt.Printf("%T \n",a)
	fmt.Printf("%d \n",a)*/
	var total float64
	for _, v := range a {
		total += v
		}
	return total/float64(len(a))
}
