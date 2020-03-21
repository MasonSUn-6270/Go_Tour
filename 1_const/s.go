package main

import (
	"fmt"
)

const a int = 11


func main()  {
	const b int = 0  /* 即使在func空间内，常量也不允许使用：= */
	fmt.Printf("the value of a is %v \n",a )
	fmt.Printf("the Type of a is %T",a )
}