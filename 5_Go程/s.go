package main

import (
	"fmt"
	"time"
)
var s string = "1"+"2"

func main()  {
	goroutine()
	s += "ss"
	fmt.Println(s)
	use_channel()
}


func say(s string)  {
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(time.Now(),s)

	}
}


func goroutine()  {  //Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到
	go say("1")
	go say("2")
	go say("3")
	go say("4")
	
	say("@")
}

/*
信道
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
*/

func add_string(s []int,channel chan int)  {
	var tar int
	for _,i := range s {
		tar += i

	}
	channel <- tar
}


func use_channel(){
	c := make(chan int)
	go add_string([]int{1,2,3,3,4,5,6},c)
	go add_string([]int{10,20,30,30,40,50,60},c)
	x,y := <-c ,<-c
	fmt.Println(x,y)
}

