package main

import "fmt"
import "time"


func collatzChainCount(n int) int{
    v:=int64(n)
    count:=1
    for ;v>1;count++{
        if v%2==0{
            v=v/2
        } else {
            v=v*3+1
        }
    }
    return count
}

func findMaxChain(max int) (n,maxchaincount int){
    n,maxchaincount=0,0
    for i:=1;i<=max;i++{
        count:=collatzChainCount(i)
        if count>maxchaincount{
            maxchaincount=count
            n=i
        }
    }
    return
}

func main(){
    fmt.Println(`Longest Collatz sequence
    The following iterative sequence is defined for the set of positive integers:
    n → n/2 (n is even)
    n → 3n + 1 (n is odd)
    Using the rule above and starting with 13, we generate the following sequence:
    13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
    It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. 
    Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
    Which starting number, under one million, produces the longest chain?
    
    考拉兹数列
    F(n)=F(n-1)/2   n是偶数时
    F(n)=F(n-1)*3+1 n是质数时
    虽然考拉兹猜想至今未被证明，但是，我们可以看到，所有给出的数列，都是1结束。
    求：找到百万以下的数，这个链最长
        `)
    starttime:=time.Now()
    fmt.Println(findMaxChain(1000000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}