package main

import "fmt"
import "time"
import "strconv"

func pow(n int)int{
    if n<=0{return 0}
    result:=1
    for i:=2;i<=n;i++{
        result*=i
    }
    return result
}


func findN(arr []int,n int)string{
    len:=len(arr)
    switch len{
        case 0:return ""
        case 1:{
            if n==1{
                return strconv.Itoa(arr[0])
            }
        }
        default:{
            if n<=pow(len){
                pow:=pow(len-1)
                d:=n/pow
                m:=n%pow
                if m==0{
                    d--
                    m=pow
                }
                return strconv.Itoa(arr[d])+findN(append(arr[:d],arr[d+1:]...),m)
            }else{
                return ""
            }
        }
    }
    return ""
}


func main(){
    starttime:=time.Now()
    fmt.Println(`
    Lexicographic permutations
    A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of 
    the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:
    012   021   102   120   201   210
    What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
    字典的排列
    排列，是有顺序的组合。例如：3124是1,2,3,4的一种排列。
    如果，一个集合包括一些字符的所有排列，我们称它为一个字典。
    请问，0,1,2,3,4,5,6,7,8,9的字典里，第1百万个排列是什么数字?
    `)
    fmt.Println(findN([]int{0,1,2,3,4,5,6,7,8,9},1000000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
}    