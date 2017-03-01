package main

import "fmt"
import "time"
import "math"


func isPrime(n int) bool{
    if n<2{return false}
    for i:=2;i<=int(math.Sqrt(float64(n))+1);i++{
        if n%i==0{return false}
    }
    return true
}

func findMaxPrimes(n int)(a,b,max int){
    maxa,maxb,maxn:=0,0,0
    for i:=-n+1;i<n;i++{
        for j:=-n;j<=n;j++{
            for k:=0;;k++{
                if !isPrime(k*k+i*k+j){break}
                if k>maxn{
                    maxa=i
                    maxb=j
                    maxn=k
                }
            }
        }
    }
    return maxa,maxb,maxn
}

func main(){
    fmt.Println(`
    Quadratic primes
    Euler discovered the remarkable quadratic formula:
        n^2+n+41
    It turns out that the formula will produce 40 primes for the consecutive integer values 0≤n≤390≤,n≤39. However, when
    n=40,402+40+41=40(40+1)+41n=40,402+40+41=40(40+1)+41 is divisible by 41,
    and certainly when n=41,412+41+41n=41,412+41+41 is clearly divisible by 41.
    Find the product of the coefficients, aa and bb, for the quadratic expression that produces the maximum number of primes for 
    consecutive values of nn, starting with n=0n=0.
    欧拉发现的这个二次方程n^2+n+41，
    是一个神奇的方程。因为对于0<=n<=39，n^2+n+41都是质数。
    另一个不可置信的二次方程是
    n^2+79n+1601
    对于0<=n<=79连续80个n，都得到了质数。
    求：在 |a|<1000,|b|≤1000的范围内，找到a,b,使
    n^2+an+b1
    得到最多的连续质数，给出a*b。
    `)
    starttime:=time.Now()
    fmt.Println(findMaxPrimes(1000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}