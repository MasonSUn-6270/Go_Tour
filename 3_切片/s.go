package main

import ("fmt"
"strings"
)



/*

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片的零值是 nil。

*/

func base()  {
	var x = []int {1,2,3,4,5,6}
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
	x= x[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
	x= x[3:5]
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
	x= x[1:2]
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}

func test_nil()  {
	var x = [5]int{}
	fmt.Printf("x len=%d cap=%d %v\n", len(x), cap(x), x)
	var y = []int{}
	fmt.Printf("y len=%d cap=%d %v\n", len(y), cap(y), y)
	var a = make([]int, 0,5) // len=0, cap=5
	fmt.Printf("a len=%d cap=%d %v\n", len(a), cap(a), a)
	var b = make([]int,5)
	fmt.Printf("b len=%d cap=%d %v\n", len(b), cap(b), b)

}

func append_()  {
	var x []int
	x = append(x,1,2,3,4,5,)
	fmt.Println(x)
	x = append(x,1,2,3,4,5,)
	fmt.Println(x)
}

func game()  {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main()  {
	base();test_nil();game();append_()
}