package main

import "fmt"

func add(x,y int)  int{
	return x+y
}

func base()  {
	i := add
	fmt.Println(i(3,5))
}


func foo(fn func(int,int) int ) int {
	return fn(5,6)
}


func main()  {
	base()
	fmt.Println(foo(add))
}