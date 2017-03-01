package main

import (
    "fmt"
    "time"
    )
    
//find the Pythagorean triplet abc that a+b+c=sum
func findTriplet(sum int) (int,int,int){
    var a,b,c int
    if sum<3 {return 0,0,0}
    for i:=1;i<=sum-2;i++{
            for j:=i+1;j<=sum-1;j++{
                a=i
                b=j-i
                c=sum-j
                if a*a+b*b==c*c{
                    return a,b,c
            }
        }
    }
    return -1,-1,-1
}

func main(){
    fmt.Println(`Special Pythagorean triplet
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
毕达哥拉斯定理
直角三角形的三条边abc,a^2 + b^2 = c^2
找出这样的abc,满足a+b+c=1000
`)
    starttime:=time.Now()
    a,b,c:=findTriplet(1000)
    fmt.Println(a,b,c)
    // 200^2= 40000
    // 375^2=140625
    // 425^2=180625
    fmt.Println("Time elapsed: ", time.Since(starttime))
}