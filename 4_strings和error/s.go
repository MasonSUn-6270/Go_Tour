package main

import "fmt"

type Name struct{name string}
/* 
type Stringer interface {
    String() string
}
*/
func (n *Name)String()  string{   //一定要命名String
	return fmt.Sprintf("我是str方法，我被调用了")
}

func __repr__()  {
	fmt.Println(&Name{"1"})
}


func (n *Name)Error()  string{  //异常处理
	return "heihei"
}

// func raise_error()  error{  //raise error
// 	return &Name{"SS"}
// }









func main()  {
	__repr__()
	
}