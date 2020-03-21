package main

import("fmt")

var a,b,c bool 
var x int =100
var y = 100
var haha,hehe = true, "hehe"


func main()  {
	walrus := "函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。"
	var i int 
	fmt.Println(a,b,c,i)
	fmt.Println(x,y)
	fmt.Println(haha,hehe)
	fmt.Println("海象操作符 : ",walrus)
}
/* 默认都为False */