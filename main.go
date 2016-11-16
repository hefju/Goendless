package main
import (
	"fmt"
	"github.com/hefju/Goendless/web"

)
func main() {


	web.Run()
    //atest()
    fmt.Println("endless")
	return
}



func atest(){

    x:=make([]int ,4)
	fmt.Printf("原地址:%p\n",x)//打印变量的地址
	x=append(x,5)
	fmt.Printf("后地址:%p\n",x)
	fmt.Println(x)
}