package main

import "fmt"

type I interface{ M() int }
type  S struct { a int ; b int }
func (s *S)M()  int{
	return s.a*s.a+s.b*s.b
}
func normal()  {
	var i I  = &S{3,4}
	/* var ii I = "hello" cannot use "hello" (type string) as type I in assignment: string does not implement I (missing M method) */
	fmt.Println(i.M())
	fmt.Printf("对象的value：%v   对象的类型 %T \n",i,i)
}

func empty_interface()  {   //空接口
	var foo interface{}
	fmt.Printf("对象的value：%v   对象的类型 %T \n",foo,foo)
	foo = 1
	fmt.Printf("对象的value：%v   对象的类型 %T \n",foo,foo)
}


func assert()  {  //断言
	var tmp interface{} = 1
	val , ok := tmp.(int)
	fmt.Println(tmp,val,ok)
}

func judge_type(x interface{})  {  //判断类型 
	switch t:=x.(type) {
	case int:
		fmt.Println("int")
		fmt.Println(t)
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("sorry  i dont know ")
	}
}

func use_judge(){
	judge_type(5)
	judge_type(5.0)
	judge_type("hello")
}

func main()  {
	normal();
	empty_interface();
	assert()
	use_judge()
}