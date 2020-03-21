package main

import "fmt"

func print(x int)  {
	fmt.Println(x)
}

func stair(x int)  {
	switch  {
	case x<10 :
		print(x+10)
	case x<100:
		print(x*x)
	default:
		print(99999999)
	}
}


func main()  {
	stair(40)
}