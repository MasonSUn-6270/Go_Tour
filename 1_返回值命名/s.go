package main

import("fmt")

func square(x,y int)  (x1,y1 int ){
	x1 = x*x
	y1 = y*y
	return
}

func main()  {
	fmt.Println(square(1,10))
}