package main

import "fmt"
import "time"
import "./lib"

//get count of cycles of 1/n
func getCycles(n int)int{
    e:=lib.Eular(n)
    // fmt.Println(eular)
    for _,i:=range(lib.GetFactors(e)){
        if lib.Pow(10,i,n)==1{
            return i
        }
    }
    return 0
}

func getMaxCycles(n int)int{
    max,maxn:=0,0
    arr:=lib.GetPrimes(n)
    for _,i:=range(arr){
        l:=getCycles(i)
        if l>max{
            max=l
            maxn=i
        }
    }
    return maxn
    /*
    This is a useful application of Fermat’s little theorem that says: 1/d has a cycle of n digits if 10^n − 1 mod d = 0 
    for prime d. It also follows that a prime number in the denominator, d, can yield up to d − 1 repeating decimal digits.
    We need to find the largest prime under 1000 that has d − 1 digits. This is called a full reptend prime. 
    这个问题的另一种描述:给定大整数n(可能是质数也可能是合数,且不知道这个数的分解形式),求最小的k使10^k ≡1 (mod n)
    对a^k mod n=1
    若n与a互素,求分母n的欧拉函数值ψ(n).那么循环节长度k必是ψ(n)的约数.
    若n与a有公因子,显然无解.
    事实上提出这个问题的初衷,是发现大数分解问题可以转化为求一个大数的倒数的循环节的长度
    给定n,在RSA加密中,n肯定是两个质数的积,设n=p*q,此时1/n的循环节的长度l|gcd(p-1,q-1),
    假定知道l的因数分解,l=l(1)^c(1）*l(2)^c(2)*...*l(k)^c(k),则l有∏[c(i)+1]个约数,将这些约数分别加上1,如果某个
    约数y(j)加一后是质数,则y(j)+1有可能是n的约数,对所有j<sqrt(n)-1的y(j)进行检验,必能找到一个恰好满足y(j)+1=
    min(p,q),这一部分所用的时间应该不会很多.于是大数问题就转化为求1/n的循环节长度l
    当然l也可能是一个很大的数,但对n为奇数的情况,l必为偶数,可以先除去所有因数2,甚至其他较小的素因子,得到l',
    然后再用相同的办法,求1/l'的循环节长度l(2)...
    即使在最坏的情况下,也有l'<n/4,因此一个300位的大整数,最多只需通过log(10^300)/log(4) <500次转换,就可以完成分解
    */
}

func main(){
    fmt.Println(`
    Reciprocal cycles
    A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:
    1/2 =  0.5 
    1/3 =  0.(3) 
    1/4 =  0.25 
    1/5 =  0.2 
    1/6 =  0.1(6) 
    1/7 =  0.(142857) 
    1/8 =  0.125 
    1/9 =  0.(1) 
    1/10 =  0.1 
    Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.
    Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
    倒数里的循环小数
    从d=1到1000里面找出一个数，使得1/d转化成的小数有最长的小数循环。
    `)
    starttime:=time.Now()
    // fmt.Println(getCycles(997))
    fmt.Println(getMaxCycles(1000))
    fmt.Println("Time elapsed: ", time.Since(starttime))
    //https://github.com/Fr0sT-Brutal/Delphi_MemoryModule?utm_medium=referral&utm_campaign=ZEEF&utm_source=https%3A%2F%2Fdelphi.zeef.com%2Fanton.frost
}    