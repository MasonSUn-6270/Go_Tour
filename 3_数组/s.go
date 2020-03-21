package main

import "fmt"

var y [2]string

var x [5]int = [5]int {1,2,3,4,5}

func main()  {
	y[0]="h";y[1]="W"
	var z = x[1:3]
	var x = [5]int{10,11,12,14,15}
	
	fmt.Println(y)
	fmt.Println(x[4],z)
	fmt.Println(x)
}

