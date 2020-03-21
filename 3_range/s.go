package main

import "fmt"

func base()  {
	var x = [5]int{100,200,300,400,500}
	for i,v := range x{
		fmt.Println(i,"&",v)
	} 
	
	/* _ */
	for _,val :=range x{
		fmt.Println(val)
	}

}

func main()  {
	base()
}