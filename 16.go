package main

import "fmt"
import "time"
import "math/big"

func digitSum(n int64) int64{
    var result int64
    value:=big.NewInt(1)
    ten:=big.NewInt(10)
    sum:=big.NewInt(0)
    x:=big.NewInt(0)
    value.Exp(big.NewInt(2),big.NewInt(n),nil)
    for value.Cmp(ten)>=0{
        x.Mod(value,ten)
        sum.Add(sum,x)
        value.Div(value,ten)
    }
    sum.Add(sum,value)
    result=sum.Int64()
    return result
}

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Power digit sum
    2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
    What is the sum of the digits of the number 2^1000?
    幂的各位之和
    2^15 = 32768，那么各位之和就是3+2+7+6+8=26
    请问：2^1000的各位之和？
    `)
    fmt.Println(digitSum(1000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}