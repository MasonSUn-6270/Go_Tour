package main

import "fmt"

func base()  {
	// var j *int 类型 *T 是指向 T 类型值的指针。其零值为 nil。
	i,foo :=  5,-99;j := &i
	fmt.Println("指针刚定义",j)
	*j ++	
	fmt.Println("指针运算后",j)
	fmt.Println(i)
	j = &foo
	*j -= 100
	fmt.Println(i,foo)
}


func main()  {
	base()
}