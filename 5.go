package main

import (
    "fmt"
    "time"
    )
var lMap=make(map[int]int)

func adjMap(a,b map[int]int){
    for k,v:=range a{
        if b[k]<v {
            b[k]=v
        }
    }
}

func findMultiple(n int){
    llMap:=make(map[int]int)
    for n>1 {
        for i:=2;i<=n;i++{
            if i==n {
                llMap[i]++
                adjMap(llMap,lMap)
                return
            }
            if n%i==0 {
                n=n/i
                llMap[i]++
                break
            }
        }
    }
    adjMap(llMap,lMap)
}

func findMiniMultiple(n int) int{
    for i:=1;i<=n;i++{
      findMultiple(i)  
    }
    result:=1
    for k,v:=range lMap{
        for v>0 {
            result*=k
            v--
        }
    }
    return result
}

func main(){
    fmt.Println(`Smallest multiple
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
2520是能被从1到10的所有数整除的最小数。请问：能被从1到20的所有数整除的最小数是多少?
        `)
    starttime:=time.Now()
    fmt.Println(findMiniMultiple(20))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}