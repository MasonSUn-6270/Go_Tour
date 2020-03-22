package main

import "fmt"

type Foo struct {x,y float64
	}

func (self Foo) Div() float64{
	return self.x/self.y
}

func method1()  {
	self := Foo{1,2}
	fmt.Println(self.Div())
}


func Div2(self Foo) float64{
	return self.x/self.y
}

func method2()  {
	self:=Foo{1,2}
	fmt.Println(Div2(self))
}

func main()  {
	method1();method2()
}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               