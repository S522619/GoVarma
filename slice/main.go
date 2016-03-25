package main

import "fmt"

func main(){
	abc := []string{
		"Sandeep", "Dilip", "Daddy", "Mommy",
	}

	for i,n:= range abc{
		fmt.Println(i,n)
	}

	bac := make([]int,3,7)
        bac[0] = 25
	bac[1] = 22
	bac[2] = 22
	fmt.Println(bac)
	fmt.Println(bac[0])
	bac[2]++

	fmt.Println(len(bac))
	fmt.Println(cap(bac))

	a := []int{1,2,3,4,5,6}


	for i, value := range a{
		fmt.Println(i,"-",value)
	}





}
