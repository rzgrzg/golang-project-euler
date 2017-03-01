package main

import "fmt"
import "time"
import "math/big"

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Factorial digit sum
    n! means n × (n − 1) × ... × 3 × 2 × 1
    For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
    and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
    Find the sum of the digits in the number 100!
    阶乘的各位之和
    `)
    value:=big.NewInt(1)
/*    for i:=1;i<=100;i++{
        value.Mul(value,big.NewInt(int64(i)))
    }*/
    value.MulRange(1,100)
    mod:=big.NewInt(0)
    ten:=big.NewInt(10)
    sum:=big.NewInt(0)
    for value.Sign()!=0{
        value.DivMod(value,ten,mod)
        sum.Add(sum,mod)
    }
    fmt.Println(value,sum)
    fmt.Println("Time elapsed: ", time.Since(starttime))
}