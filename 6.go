package main

import (
    "fmt"
    "time"
    )
    
func calc(n int) int{
    /*
    (1+2+...+n)^2=[n(n+1)/2]^2
    n^2=[(n+2)(n+1)n-(n+1)n(n-1)]/3-n
    
    1^2=[3*2*1-2*1*0]/3-1
    2^2=[4*3*2-3*2*1]/3-2
    3^2=[5*4*3-4*3*2]/3-3
    ...
    n^2=[(n+2)(n+1)n-(n+1)n(n-1)]/3-n
    total:
    =[(n+2)(n+1)n]/3-n(n+1)/2
    */
    return n*(n+1)/2*n*(n+1)/2-(n+2)*(n+1)*n/3+n*(n+1)/2
}
    
func main(){
    fmt.Println(`Sum square difference
The sum of the squares of the first ten natural numbers is,
1^2 + 2^2 + ... + 10^2 = 385
The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)^2 = 55^2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
自然数1-10的平方和=385
自然数1-10的和的平方=3025
自然数1-10的和的平方比平方和大2640
请问：自然数1-100的和的平方比平方和大多少
        `)
    starttime:=time.Now()
    fmt.Println(calc(100))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}