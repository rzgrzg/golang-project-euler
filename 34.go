package main

import (
    "fmt"
    "time"
    )
    
/*
分析：
先得确定一个上限。
9!                  =362880
9!+9!+9!+9!         =1451520
9!+9!+9!+9!+9!      =1814400
9!+9!+9!+9!+9!+9!   =2177280
9!+9!+9!+9!+9!+9!+9!=2540160
可见，最大可能的数字可以是2540160
*/

func isDigitfactorials(n int) bool {
    sum:=0
    m:=n
    for m>0{
        digit:=m % 10
        m/=10
    }
    return n==sum
}

func sumDigit() int {
    const max=2540160
    sum:=0
    for i:=3;i<max;i++{
        if isDigitfactorials(i){
            sum+=i
        }
    }
    return sum
}


func main(){
    fmt.Println(`
    阶乘数字
    145=1! + 4! + 5!
    找出所有的数字，每一个都等于各位的阶乘。
    计算所有这些数字的和。
    (1!=1和2!=2不算)
    `)
    starttime:=time.Now()
    fmt.Println(sumDigit())
    fmt.Println("Time elapsed: ", time.Since(starttime))    
}