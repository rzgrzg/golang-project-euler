package main

import "fmt"
import "time"

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Cake Icing Puzzle
    
    Adam plays the following game with his birthday cake.
    He cuts a piece forming a circular sector of 60 degrees and flips the piece upside down, with the icing on the bottom.
    Amazingly, this works for any piece size, even if the cutting angle is an irrational number: 
    all the icing will be back on top after a finite number of steps.
    蛋糕结冰难题
    亚当用他生日的蛋糕做了个游戏。
    蛋糕的上面有冰。
    他在蛋糕上60度弧形的一块，翻过来。然后把蛋糕逆时针转60度。
    重复了12次，结果冰又全在上面了。
    这个没什么，
    重要的是，亚当又有了个惊人的发现。
    上面这个流程可以在任意角度下实现，即使是无理数，也能够在有限的次数下，让冰全部在蛋糕的上面。
    亚当又换了个方式。
    
    
    
    `)
    fmt.Println("Time elapsed: ", time.Since(starttime))
}