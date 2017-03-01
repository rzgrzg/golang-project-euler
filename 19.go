package main

import "fmt"
import "time"

type
    Weekday int
const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (d Weekday)String()string{
    switch d{
        case Sunday:{return"Sunday"}
        case Monday:{return "Monday"}
        case Tuesday:{return "Tuesday"}
        case Wednesday:{return "Wednesday"}
        case Thursday:{return "Thursday"}
        case Friday:{return "Friday"}
        case Saturday:{return "Saturday"}
        default:return ""
    }
}

func isLeapYear(year int)bool{
    if year%4!=0{
        return false
    }else{
        if year%100==0{
            if year%400==0{
                return true
            }else{
                return false
            }
        }else{
            return true
        }
    }
}

func getWeekDay(year,month,day int) Weekday{
    days:=1
    if year>1900{
        for i:=1900;i<year;i++{
            if isLeapYear(i){
                days+=366
            } else {
                days+=365
            }
        }
    }else if year<1900{
        for i:=year;i<=1900;i++{
            if isLeapYear(i){
                days-=366
            } else {
                days-=365
            }
        }
    }
    if month>1{
        for i:=1;i<month;i++{
            switch i{
                case 2:{
                    if isLeapYear(year){
                        days+=29
                    }else{
                        days+=28
                    }
                }
                case 4,6,9,11:{days+=30}
                default:{days+=31}
            }
        }
    }
    if day>1{
        days+=day-1
    }
    days=days%7
    if days<0{days+=7}
    return Weekday(days)
}

func countSundays() (count int) {
    for i:=1901;i<2001;i++{
        for j:=1;j<=12;j++{
            if getWeekDay(i,j,1)==Sunday{
                count++
            }
        }
    }
    return count
}

func main(){
    starttime:=time.Now()
    fmt.Println(`
    Counting Sundays
    You are given the following information, but you may prefer to do some research for yourself.Sundays
    1 Jan 1900 was a Monday.
    Thirty days has September,
    April, June and November.
    All the rest have thirty-one,
    Saving February alone,
    Which has twenty-eight, rain or shine.
    And on leap years, twenty-nine.
    A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
    
    How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
    星期天的数量
    1900年1月1日是星期一
    请计算，20世纪的(从1901年1月1日到2000年12月31日)，所有的每月1号一共有多少个星期天
    `)
    // fmt.Println(getWeekDay(2010,1,1))
    fmt.Println(countSundays())
    fmt.Println("Time elapsed: ", time.Since(starttime))
}