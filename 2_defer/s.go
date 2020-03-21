package main

import "fmt"

func base()  {
	defer fmt.Println("hello")
	defer fmt.Println("world")
	fmt.Println("我系渣渣辉")
}


func test_stack(){


	var i int = 1	
	for ; i <100 ;i++{
		
		defer fmt.Println(i)
	}
	fmt.Println("hi",i)


}
func main()  {
	base()
	test_stack()
}