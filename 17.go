package main

import "fmt"
import "time"
import "strings"

var (
    thousands=[]string{"thousand","million","billion"}
    decades=[]string { "","","twenty","thirty","forty","fifty","sixty","seventy","eighty","ninety","hundred" }
    unit=[]string{ "zero","one","two","three","four","five","six","seven","eight","nine","ten","eleven","twelve","thirteen","fourteen","fifteen","sixteen","seventeen","eighteen","nineteen" }
)

func N2E(n int64) string{
    if n==0{return unit[0]}
    if n<0{
        n=-n
        return "Minus "+N2E(n)
    }
    //thousands
    if n>=1000000000{
        pref:=n/1000000000
        n=n%1000000000
        if n==0{
            return N2E(pref)+" "+thousands[2]
        } else {
            return N2E(pref)+" "+thousands[2]+" "+N2E(n)
        }
    }
    if n>=1000000{
        pref:=n/1000000
        n=n%1000000
        if n==0{
            return N2E(pref)+" "+thousands[1]
        } else {
            return N2E(pref)+" "+thousands[1]+" "+N2E(n)
        }
    }
    if n>=1000{
        pref:=n/1000
        n=n%1000
        if n==0{
            return N2E(pref)+" "+thousands[0]
        } else {
            return N2E(pref)+" "+thousands[0]+" "+N2E(n)
        }
    }
    if n>=100{
        pref:=n/100
        n=n%100
        if n==0{
            return N2E(pref)+" "+decades[10]
        } else {
            return N2E(pref)+" "+decades[10]+" and "+N2E(n)
        }
    }

    if n>=20{
        pref:=n/10
        n=n%10
        if n==0{
            return decades[pref]
        } else {
            return decades[pref]+"-"+N2E(n)
        }
    } else {
        return unit[n]
    }
}

func countNumber(max int)int{
    count:=0
    for i:=1;i<=max;i++{
        str:=N2E(int64(i))
        count+=strings.Count(str,"")-strings.Count(str," ")-strings.Count(str,"-")-1
    }
    return count
}

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Number letter counts
    If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
    If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used? 
    NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains
    20 letters. The use of "and" when writing out numbers is in compliance with British usage.
    不认识英语的请直接下一题。
    `)
    fmt.Println(countNumber(1000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}