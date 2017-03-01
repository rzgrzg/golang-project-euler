package main

import (
    "fmt"
    "time"
    "math"
    )

func isPrime(n int) bool{
    // for i:=2;i<=n/2;i++{
    for i:=2;i<=int(math.Sqrt(float64(n))+1);i++{
        if n%i==0{return false}
    }
    return true
}

func main(){
    fmt.Println(`10001st prime
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
What is the 10 001st prime number?
质数的前6个是2,3,5,7,11,13。请问：第10 001个质数是多少?
        `)
    starttime:=time.Now()
    count:=1
    for i:=3;;i+=2{
        if isPrime(i){
            count++
            if count==10001{
                fmt.Println(count,i)
                break
            }
        }
    }
    
    
    fmt.Println("Time elapsed: ", time.Since(starttime))
}