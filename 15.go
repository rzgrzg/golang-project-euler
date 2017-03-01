package main

import "fmt"
import "time"

//阶乘
func fact(n int)int64{
    if n<0{return -1}
    product:=int64(1)
    for i:=1;i<=n;i++{
        product*=int64(i)
    }
    return product
}

func findRoutes(n int)int64{
    // 实际上这个问题，就是个x=n,y=n的组合问题
    //结果是：2n的排列，除以(n的排列*2)
    //对于2*2方框来说，就是4*3*2*1/2/2=6
    //int64位数不够啊啊啊
    return fact(n*2)/fact(n)/fact(n)
}
func main(){
    starttime:=time.Now()
    fmt.Println(`
    Lattice paths
    Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, 
    there are exactly 6 routes to the bottom right corner.
    How many such routes are there through a 20×20 grid?
    方框里的路径
    2*2的方框，只能往下和往右走，一共有6种走法，可以达到右下角
    请问：20*20的方框，有多少种走法？
    `)
    fmt.Println(fact(40))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}