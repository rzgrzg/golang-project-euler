package main

import "fmt"
import "time"

func pow(a,b int)int{
    if b<0{return -1}
    if b==0{return 0}
    if b==1{return a}
    result:=1
    for i:=0;i<b;i++{
        result*=a
    }
    return result
}

//return the posibly max A where Digits of A,(a,b,c,d,e,...) a^n+b^n+...==A
func findMax(n int)int{
    result:=0
    base:=pow(9,n)
    /*
    9:6561
    99:13122
    999:19683
    9999:26244
    99999:32805
    */
    sum:=0
    for i:=1;;i++{
        sum+=pow(10,i-1)*9
        result=i*base
        if sum>result{
            break
        }
    }    
    return result
}

func sumOfN(n int)int{
    result:=0
    var sum int
    for i:=2;i<=findMax(n);i++{
        value:=i
        sum=0
        for value>0{
            sum+=pow(value%10,n)
            value/=10
        }
        if sum==i{
            result+=sum
        }
    }
    return result
}

func main(){
    fmt.Println(`
    Digit fifth powers
    Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:
1634 = 1^4 + 6^4 + 3^4 + 4^4
 8208 = 8^4 + 2^4 + 0^4 + 8^4
 9474 = 9^4 + 4^4 + 7^4 + 4^4
As 1 = 1^4 is not a sum it is not included.
The sum of these numbers is 1634 + 8208 + 9474 = 19316.
Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
    5次幂的和
    上面的3个4位数，各位的4次幂之和，刚好和自身相等。
    那么，所有符合这个条件的数的和，是19316
    请问，对于5次幂，这个和是多少？
    `)
    starttime:=time.Now()
    fmt.Println(sumOfN(5))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}