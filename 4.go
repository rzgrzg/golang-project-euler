package main

import (
    "fmt"
    "time"
    )
    
func ispalindromic(num int) bool{
    source:= num
    tnum:= 0
    for num != 0 {
        tnum = tnum*10 + num%10
        num = num / 10
    }
    if tnum == source {
        return true
    }    
    return false
}

func findpalindromic() int{
    result:=0
    for i:=999;i>900;i--{
        for j:=999;j>i;j--{
            value:=i*j
            if ispalindromic(value) {
                if value>result {result=value}
            }
        }
    }
    return result
}
    
func main(){
    fmt.Println(`Largest palindrome product
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
Find the largest palindrome made from the product of two 3-digit numbers.
回文数，指的是从左向右读和从右向左读完全相同的数。2位数相乘所得到的最大回文数是9009。
找出3位数相乘所得的最大回文数。
        `)
    starttime:=time.Now()
    fmt.Println(findpalindromic())
    fmt.Println("Time elapsed: ", time.Since(starttime))    
}
    
