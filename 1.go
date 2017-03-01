package main

import "fmt"

func isDiv(n int) bool {
    if n<=0||n%3==0 || n%5==0 {
        return true
    }
    return false
}

func main(){
    if isDiv(5) {
        fmt.Println("5")
    }
    sum:=0
    for i := 0; i < 1000; i++{
        if isDiv(i) {
            sum+=i
        }
    }
    fmt.Println("total is ",sum)
    
}