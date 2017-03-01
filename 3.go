package main

import "fmt"
import "time"

func findprime(n uint64) uint64{
    for n>1 {
        for i:=uint64(2);i<=n;i++{
            if i==n {return n}
            if n%i==0 {
                n=n/i
                break
            }
        }
    }    
    return n
}

func main(){
    fmt.Println(`
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
        `)
    starttime:=time.Now()
    // fmt.Println(findprime(13195))
    fmt.Println(findprime(600851475143))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}