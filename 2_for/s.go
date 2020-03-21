package main

import "fmt"

func base()  {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func without_init_finish()  {
	i := 0
	for  ; i < 50;  i++{
		i += i
		
	}
	fmt.Println(i)
}

//无线循环
func inifinty()  {
	for  {
		fmt.Println("inf")
	}
}

func main()  {
	base()
	without_init_finish()
	inifinty()
}