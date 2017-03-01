package main

import "fmt"
import "time"
import "math/big"

func findN(n int)int{
    a:=big.NewInt(1)
    b:=big.NewInt(1)
    c:=big.NewInt(1)
    max:=big.NewInt(1)
    count:=2
    max.Exp(big.NewInt(10),big.NewInt(int64(n-1)),nil)
    for{
        //a=b,b=a+b
        c.Set(b)
        b.Add(a,b)
        a.Set(c)
        count++
        if b.Cmp(max)==1{
            break
        }
    }
    return count
}

func main(){
    fmt.Println(`
    1000-digit Fibonacci number
    The Fibonacci sequence is defined by the recurrence relation:
    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
    Hence the first 12 terms will be:
    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144
    The 12th term, F12, is the first term to contain three digits.
    What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
    斐波那契数字，1000位
    斐波那契函数是这样一个函数:
    F1 = 1
    F2 = 1
    Fn = Fn−1 + Fn−2
    F12=144是第1个3位数的斐波那契数字。 
    请问：第1个1000位数的斐波那契数字是第几个数字？
    `)
    starttime:=time.Now()
    fmt.Println(findN(1000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}    