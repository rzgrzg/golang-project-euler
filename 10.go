package main

import (
    "fmt"
    "time"
    "math"
    )
    
func isPrime(n int) bool{
    if n<2{return false}
    for i:=2;i<=int(math.Sqrt(float64(n))+1);i++{
        if n%i==0{return false}
    }
    return true
}

func findPrimeSum(n int) int64{
    if n<2 {return -1}
    var result int64=2
    for i:=3;i<n;i+=2{
        if isPrime(i){
            result+=int64(i)
        }
    }
    return result
}

func main(){
    fmt.Println(`Summation of primes
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
Find the sum of all the primes below two million.
质数之和
小于10的质数2,3,5,7的和是17
请问：小于两百万的指数之和是多少？
`)
    starttime:=time.Now()
    fmt.Println(findPrimeSum(2000000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}