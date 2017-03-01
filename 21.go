package main

import "fmt"
import "time"

const mod=9901

func fast_pow(a int,n int) int{
    result := 1
    pre := a
    a%=mod
    for n>0{
        if(n&1)!=0{
            result *= pre
            result %= mod
        }
        pre*=pre
        pre%=mod
        n>>=1
    }
    return result
}

//二分求所有因子的和a^0+...+a^n
func toCalculate(a int,n int )int{
    result := 1
    if(n==0){return 1}
    if(n%2!=0){
        result = ((1+fast_pow(a,n/2+1))%mod)*toCalculate(a,n/2)
    } else {
        result = ((1+fast_pow(a,n/2+1))%mod)*toCalculate(a,n/2-1)+fast_pow(a,n/2)
    }
    result%=mod;
    return result
}
/*
求一个数的所有因数之和 = （a^0+a^1+a^2+a^3+….a^n） * （ a1^0+a1^1+a1^2+a1^3+….a1^n1）* （a2^0+a2^1+a2^2+a2^3+….a2^n2）…..*（an^0+an^1+an^2+an^3+….an^nn） 
 其中a,a1,a2,,,,代表它的质因数，n,n1,n2….nn代表它所代表的质因数的个数*/
 
func sumDivisorsOfMap(lMap map [int]int) (sum int){
    sum=1
    for k,v:=range lMap{
        sum*=toCalculate(k,v)
    }
    return
}

func sumDivisors(n int) int{
    l:=n
    lMap:=make(map [int]int)
    for n>1 {
        for i:=2;i<=n;i++{
            if i==n {
                lMap[i]++
                return sumDivisorsOfMap(lMap)-l//because include n self in sum
            }
            if n%i==0 {
                n=n/i
                lMap[i]++
                break
            }
        }
    }    
    return 0
}

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Amicable numbers
    Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
    If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
    For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. 
    The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
    Evaluate the sum of all the amicable numbers under 10000.
    友好数字
    设，d(n)为所有n的因数之和。
    如果，存在两个数a和b,使得d(a)=b并且d(b)=a,a和b为友好数字。
    求：所有10000以下的所有友好数字的和。
    `)
    sum:=0
    for i:=1;i<10000;i++{
        a:=sumDivisors(i)
        b:=sumDivisors(a)
        if a<b&&b==i{
            fmt.Println(a,b)
            sum+=a
            sum+=b
        }
    }
    fmt.Println(sum)
    fmt.Println("Time elapsed: ", time.Since(starttime))
}