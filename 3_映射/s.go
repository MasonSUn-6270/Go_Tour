package main

import "fmt"

func base()  {
	student := map[int]string{
		1:"Sam",
		2:"Ben",
	}
	fmt.Println("student",student)
	student1 := student[1]
	student3 := student[3]  // 
	student4_,ok := student[4]  // ,ok 返回状态
	fmt.Println(1,student1,3,student3)
	fmt.Println("student4",student4_,ok)
}

func main()  {
	base()
}