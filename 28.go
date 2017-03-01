package main

import "fmt"
import "time"

func sumOfDiagonals()int{
    const n=1001
    if n%2==0{return 0}
    x:=n/2
    y:=n/2
    m:=n/2
    dir:=0
    radiu:=0
    var slice [n][n] int
    for i:=1;i<=n*n;i++{
        slice[x][y]=i
        switch (dir % 4){
            case 0:{
                y++
                if y-m>=radiu+1{
                    dir++
                    radiu++
                }
            }
            case 1:{
                x++
                if x-m>=radiu{
                    dir++
                }
            }
            case 2:{
                y--
                if m-y>=radiu{
                    dir++
                }
            }
            case 3:{
                x--
                if m-x>=radiu{
                    dir++
                }
            }
        }
    }
    sum:=-1
    for i:=0;i<n;i++{
        sum+=slice[i][i]
        sum+=slice[i][n-i-1]
    }
    return sum
}

func main(){
    fmt.Println(`
    Number spiral diagonals
    Starting with the number 1 and moving to the right in a clockwise direction a 5 by 5 spiral is formed as follows:
    It can be verified that the sum of the numbers on the diagonals is 101.
    What is the sum of the numbers on the diagonals in a 1001 by 1001 spiral formed in the same way?
    
    数字螺旋
    21 22 23 24 25
    20  7  8  9 10
    19  6  1  2 11
    18  5  4  3 12
    17 16 15 14 13
    向右，顺时针，展开这样一个表，我们发现，两条对角线上各数字之和，是101。
    请问：1001*1001的这样一个表，对角线上数字之和是多少？
    `)
    starttime:=time.Now()
    fmt.Println(sumOfDiagonals())
    fmt.Println("Time elapsed: ", time.Since(starttime))
}