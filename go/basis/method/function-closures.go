/*************************************************************************
	> File Name: function-closures.go
	> Author: small_ant 
	> Mail: xms.chnb@gmail.com 
	> Created Time: 2020年12月02日 星期三 11时43分56秒
    > Description: 函数的闭包
 ************************************************************************/
package main
import "fmt"
func adder() func(int) int{
    sum := 0
    return func(x int)int{
        sum += x
        return sum
    }
}
// 函数adder返回一个闭包， 每个闭包都被绑定在其各自的sum变量上

func main(){
    pos, neg := adder(),adder()
    for i := 0; i <10; i++{
        fmt.Println(pos(i),neg(-2*i))
    }
}
