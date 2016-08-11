package main
import "fmt"
func main() {
    atest()
    fmt.Println("endless")
}
func atest(){

    x:=make([]int ,4)
    fmt.Printf("原地址:%p\n",x)
    x=append(x,5)
    fmt.Printf("后地址:%p\n",x)
    fmt.Println(x)
}