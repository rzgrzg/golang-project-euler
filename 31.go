package main

import (
    "fmt"
    "time"
    "sort"
    )
    
func ggg(mount int,base []int)[]int{
    // result:=make([]int,0)
    resultList:=[]int{0}
    sort.Ints(base)
    /*
        #目标200  
    target = 200  
    coinList = [1,2,5,10,20,50,100,200]  
    #准备记录过程中的所有线路数量  
    resultList=[0] * (target+1)  
    #第零个是1  
    resultList[0]=1  
    for i in range(0,len(coinList)):  
        #测试每一个类型  
        for j in range(coinList[i],target+1):  
            #最新的线路数量，更新  
            resultList[j] += resultList[j - coinList[i]]  
    return resultList[200]  
    */
    
    return resultList
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
    fmt.Println(ggg(200,[]int{1,2,5,10,20,50,100}))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}