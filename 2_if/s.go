package main

import "fmt"

func zero_bigger_samller(x int) string {
	if x>0 {
		return ">0"
	} else{   /*if…else 语句中的 else 必须和 if 的 ’ } ’ 在同一行，否则编译错误*/
		return "<=0"
	}

	
}

func simpify(x int ) string {
	if y := -x; y<0 {
		return "x>0"
	} else {
		fmt.Println("我都被你搞晕了")
	}
	return "x<=0"
}


func main()  {
	fmt.Println(zero_bigger_samller(10))
	fmt.Println(simpify(-4))
}