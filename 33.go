package main

import (
    "fmt"
    "time"
    )

    
func enumFrac()[]int{
    result:=[]int{}
    for a:=0;a<10;a++{
        for b:=a+1;b<10;b++{
            if (a==0)||(b==0){continue}
            for c:=0;c<10;c++{
                if 9*a*b+b*c==10*a*c{
                    result=append(result,10*a+c,10*c+b)
                }
                if 9*a*b+a*c==10*b*c{
                    result=append(result,10*c+a,10*b+c)
                } 
            }
        }
    }
    return result
}

func gcd(x, y int) int {
	tmp := x % y
	if tmp > 0 {
		return gcd(y, tmp)
	} else {
		return y
	}
}

func sumFrac()int{
    arr:=enumFrac()
    a,b:=1,1
    for i:=range(arr){
        if i%2==0{
            a*=arr[i]
        } else {
            b*=arr[i]
        }
    }
    return b/gcd(a,b)
}

func main(){
    fmt.Println(`
    找出具有一种奇怪约分性质的所有分数
     49/98是个奇怪的分数，很多人会把49/98简化成49/98 = 4/8，是的，49/98分子和分母都把9去掉就成了4/8。
     同样，30/50 = 3/5，把0去掉就行。但这种在去掉的数在同位置的不算。
     只统计ac/cb=a/b这种，总共有
     4个这种分数，小于1,并且分子和分母都是两位数。
     求这四个分数的乘积，最简后的分母？
    `)
    starttime:=time.Now()
    fmt.Println(sumFrac())
    fmt.Println("Time elapsed: ", time.Since(starttime))    
}
