package main

import "fmt"

type int_list struct{
	x int ;y int 
}


func base()  {
	fmt.Println("base",int_list{0,0})
}

func attrs()  {
	var v int_list = int_list{1,10}
	v.x = -1
	fmt.Println("the value of x is ",v.x)
}

func invisible_declare(){
	lazy := int_list{}
	lazy1 := int_list{y:99}
	p := &int_list{1,2}
	fmt.Println(lazy,lazy1,p)
}

func main()  {
	base() ; attrs();invisible_declare()
}