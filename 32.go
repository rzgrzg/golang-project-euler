package main

import (
    "fmt"
    "time"
    )
    
func main(){
    fmt.Println(`
Pandigital products
We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.
The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.
Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.
HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

Pandigital 
独数,指的是这样一种数字，从1到n(n<=9)都出现了一次，也只出现了一次。
7254是一个神奇的数字，应为39 × 186 = 7254，而39,186,7254三个数字凑齐了一个独数。
请问：
所有像7254一样能和他的因子凑齐一个1-9独数的书，他们的和是多少？
    `)
    starttime:=time.Now()
    fmt.Println("Time elapsed: ", time.Since(starttime))
}