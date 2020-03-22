package main

import (
	"fmt"

)


/* go 没有类 */

type Dog struct {
	foot_num int ; eyes_num int ; name string
}


func (self Dog)call()  {
	fmt.Println(self.name)
}

/*  果然用变量做接受者呢 */
/* 不可以 */
type A float64

func (a A)Double() int {
	return int(a*2)
}

func (d *Dog)Trible(){
	d.foot_num = d.foot_num * 3
	d.eyes_num = d.eyes_num * 3
	 
}
func (d *Dog)Add() int{
	return d.foot_num+d.eyes_num
}
type Api interface{
	Add() int ;

}


var sdsd Dog  

func main()  {
	sam := Dog{4,2,"Sam"}
	sam.Trible()
	fmt.Println(sam)
	sam.call()
	
	kk := A(2)
	fmt.Println(kk.Double())
	var a Api = &sam
	fmt.Println("孙连城万岁！",a.Add())
}
