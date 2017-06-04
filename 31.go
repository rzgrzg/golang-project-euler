package main

import (
    "fmt"
    "time"
    "sort"
    )
    
func ggg(mount int,coinList []int)int{
    if mount<=0 {return 1}
    sort.Ints(coinList)
    resultList:=make([]int,mount+1)
    resultList[0]=1
    for i:=range(coinList){
        for j:=coinList[i];j<=mount;j++{
            resultList[j] += resultList[j - coinList[i]]
        }
    }
    return resultList[mount]  
    

    /* define T(t,c) To the count of results
    T(t,c)=1 where t=0 or c=1
    T(t,c)=T(t,s(c))+T(t-c,s(c))+T(t-2*c,s(c))+...+T(t-n*c,s(c)) n=0..t/c
    s(c)=the value of the next coin smaller than c
    T(t,c)=count of ways to make amount with coid c or smaller
    
    more advance
    T(t,c)=1              t=0 or c=1
    T(t,c)=T(t,s(c))      t<c
    T(t,c)=T(t,s(c))+T(t-c,c)
    */
}
    
func main(){
    fmt.Println(`
    Coin sums
In England the currency is made up of pound, £, and pence, p, and there are eight coins in general circulation:
1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
It is possible to make £2 in the following way:
1×£1 + 1×50p + 2×20p + 1×5p + 1×2p + 3×1p
How many different ways can £2 be made using any number of coins?    
    有面值分别为1，2，5，10，20，50，100，200的硬币，现要凑齐200的面值，请问一共有多少种凑法。
    `)
    starttime:=time.Now()
    fmt.Println(ggg(200,[]int{1,2,5,10,20,50,100,200}))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}