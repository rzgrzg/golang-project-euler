package main

import (
    "fmt"
    "time"
    )

func countDivisors(n int) (count int){
    count=1
    lMap:=make(map [int]int)
    for n>1 {
        for i:=2;i<=n;i++{
            if i==n {
                lMap[i]++
                for _,v:=range(lMap){
                    count*=v+1
                }
                count--
                return
            }
            if n%i==0 {
                n=n/i
                lMap[i]++
                break
            }
        }
    }    
    return
}

func triangular(n int) int{
    return n*(n+1)/2
}

func findTriangularNumber(divisorsCount int) int{
    for i:=1;;i++{
        value:=triangular(i)
        count:=countDivisors(value)
        if count>=divisorsCount{
            return value
        }
    }
    return 0
}

func main(){
    fmt.Println(`Highly divisible triangular number
    The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 
    + 7 = 28. The first ten terms would be:
    1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...
    Let us list the factors of the first seven triangle numbers:
     1: 1
     3: 1,3
     6: 1,2,3,6
    10: 1,2,5,10
    15: 1,3,5,15
    21: 1,3,7,21
    28: 1,2,4,7,14,28
    We can see that 28 is the first triangle number to have over five divisors.
    What is the value of the first triangle number to have over five hundred divisors?
    三角中的因数
    1      = 1
    12     = 3
    123    = 6
    1234   =10
    12345  =15
    123456 =21
    1234567=28
    在上面的数列中，28有5个因数
    请问：第一个拥有500个因数的数是多少？
    `)
    
    starttime:=time.Now()
    fmt.Println(findTriangularNumber(500))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}