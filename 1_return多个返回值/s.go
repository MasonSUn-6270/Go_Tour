package main

import(
	"fmt"
	"math/rand"
)

func double_rand()  (int,int){
	return rand.Intn(10),rand.Intn(20)
}

func main()  {
	fmt.Println(double_rand())
}