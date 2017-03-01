package main

import "fmt"
import "time"
import "io/ioutil"
import "sort"
import "strings"

func nameScores(filename string) int{
    b, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Print(err)
    }
    text:=strings.Split(string(b),",")
    sort.Strings(text)
    scores:=0
    for k,v:=range(text){
        score:=0
        for _,char := range strings.Trim(v,`"`){
            score+=int(char-'A'+1)
        } 
        scores+=score*(k+1)
    }
    return scores
}

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Names scores
    Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, 
    begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value 
    by its alphabetical position in the list to obtain a name score.
    For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, 
    is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.
    What is the total of all the name scores in the file?
    名字的积分
    p022_names.txt文件里有很多的名字（超过5000），如果把它们排序，再用顺序值（a=1,c=3...不管大小写）相加，乘以在列表中的顺序值，
    那么，整个文件里所有的名字的积分和是多少？
    `)
   fmt.Println(nameScores("p022_names.txt"))
   fmt.Println("Time elapsed: ", time.Since(starttime))
}