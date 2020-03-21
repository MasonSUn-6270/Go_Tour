package main
import(
	"fmt"
)

func add(x int ,y int ) int {
	return x + y
}
/*  形参可以一起申明类型
func add(x,y int)  int {
	...
}
*/
func main()  {
	fmt.Println(add(1,2))
}