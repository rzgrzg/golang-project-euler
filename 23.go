package main

import "fmt"
import "time"



func fast_pow(a int,n int) int{
    var result int= 1
    for i:=0;i<n;i++{
        result*=a
    }
    return result
}

//求所有因子的和a^0+...+a^n,等差数列求和
func toCalculate(a int,n int)int{
    if a==1{
        return a*(n+1)
    } else {
        return (fast_pow(a,n)*a-1)/(a-1)
    }
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

func nonAbundantSums(limit int)int{
    sums:=0
    abundants:=make(map[int]bool)
    abundantsums:=make(map[int]bool)
    //prepare abundant list
    for i:=1;i<limit;i++{
        sum:=sumDivisors(i)
        if sum>i{
            abundants[i]=true
        }
    }
    for k1,v1:=range(abundants){
        if v1{
            for k2,v2:=range(abundants){
                if v2{
                     abundantsums[k1+k2]=true
                }
            }    
        }
    }
    for i:=1;i<=limit;i++{
        if !abundantsums[i]{
            sums+=i
        }
    }
    return sums
}
    


func main(){
    starttime:=time.Now()
    fmt.Println(`
    Non-abundant sums
        A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. 
    For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.
    A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.
        As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two 
    abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum 
    of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the 
    greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.
    Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
    富裕数的问题
    一个完美的数字，指的是，它的因数之和和它相等。
    列如： 1 + 2 + 4 + 7 + 14 = 28
    如果，n的因数之和<n,那n被称为不富裕数
    如果，n的因数之和>n,那n被称为富裕数
    研究表明，所有大于28123的数字，都可以表示为两个富裕数之和。
    计算出所有不能表示为两个富裕数之和的正整数之和。
    `)
    fmt.Println(nonAbundantSums(28123))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}    