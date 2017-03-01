package lib

/*
1.(a + b) % p = (a % p + b % p) % p
2.(a - b) % p = (a % p - b % p) % p
3.(a * b) % p = (a % p * b % p) % p
4. a ^ b % p = ((a % p)^b) % p
*/

/*数论四大定理：
1.费马小定理
    对于素数 M 任意不是 M 的倍数的 b，都有：b ^ (M-1) = 1 (mod M)
    用途1，测试一个数是否质数
    如果对于任意满足1 < b < p的b下式都成立：
    b^(p-1)≡1(mod p),则p必定是一个质数。
    实际上，没有必要测试所有的小于p的自然数，只要测试所有的小于p的质数就可以了。
    */
/*    
2.欧拉定理(欧拉函数)
    φ(x)=小于等于x的数中与x互质的数的数目
    φ(x)=x(1-1/p1)(1-1/p2)...(1-1/pN)//p1...pN指x的所有质因数，重复的算一个
    如：12=2*2*3那么φ（12）=12*（1-1/2）*(1-1/3)=4,4个，就是1,5,7,11
    引申：欧拉函数是积性函数——若m,n互质， φ(mn)=φ(m)φ(n)
    当n为奇数时，φ(2n)=φ(n)
    若n为质数则φ(n)=n-1
    引申:
    a^φ(n)=1 (mod n)//当n为质数时等同于费马小定理，a和n必须互质。
    RSA算法中也有用到了eula函数
    de=1 mod φ(n)是计算机安全学中的加密算法RSA， RSA算法中de=1modφ(n)表示de与1关于φ(n)同余，也就是说1除以φ(n)的余数与1除以de的余数相同。
　　例如：p=3,q=11,d=7;φ(n)=(p-1)(q-1); 
    */
func Eular(n int)int{
    /*线性筛O(n)时间复杂度内筛出maxn内欧拉函数值*/
    ret:=1
    for i:=2;i*i<=n;i++{
       if n%i==0{
            n/=i
            ret*=i-1
            for n%i==0{
                n/=i
                ret*=i
            }
        }
    }
    if n>1 {ret*=n-1}
    return ret
}

func GetPrimes(n int)[]int{
    if n<0{return nil}
    prime:=make([]int,n,n)
    v:=make(map[int]bool)
    size:=0
    for i:=2;i<n;i++{
        if !v[i]{
            size++
            prime[size]=i
        }
        j:=1
        for ((j<=size)&&(prime[j]*i<n)){
            v[i*prime[j]]=true
            if i%prime[j]==0 {break}
            j++
        }
    }
    return prime[1:size+1]
}

func GetFactors(n int)[]int{
    result:=make([]int,0,n)
    for i:=2;i<=n;i++{
        if n%i==0{
            result=append(result,i)
        }
    }
    return result
}

// expNNMontgomery calculates x**y mod m using Montgomery representation.
//蒙哥马利约减
//原理其实很好懂，假设要计算A^B，即底数是A，指数是B。把B写成二进制形式，拿4位来举例：B＝b4 b3 b2 b1（二进制）。
//先用B＝1111（二进制）来做解释。显然
//A^2 ＝ A × A， 
//A^4 ＝ A^2 × A^2， 
//A^8 ＝ A^4 × A^4。 
//假如B的二进制位数更多，则依此类推。 
func expNNMontgomery(x, y, m int) int {
    res:=1
    for y!=0{
        if (y&1)!=0{
            res = (res*x) % m
        }
        y >>= 1
        x = (x*x) % m  
    }
    return res
}

func expNNWindowed(x, y, m int) int {
/*    const n=4
    var zz, r int
    var powers [1 << n]int
    powers[0] = 1
    powers[1] = x
	for i := 2; i < 1<<n; i += 2 {
		p2, p, p1 := &powers[i/2], &powers[i], &powers[i+1]
		*p = (*p2)*(*p2)
		zz = r / *p
		r = 
		zz, r = zz.div(r, *p, m)
		*p, r = r, *p
		*p1 = p1.mul(*p, x)
		zz, r = zz.div(r, *p1, m)
		*p1, r = r, *p1
	}    
	

	z = z.setWord(1)

	for i := len(y) - 1; i >= 0; i-- {
		yi := y[i]
		for j := 0; j < _W; j += n {
			if i != len(y)-1 || j != 0 {
				// Unrolled loop for significant performance
				// gain.  Use go test -bench=".*" in crypto/rsa
				// to check performance before making changes.
				zz = zz.mul(z, z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z

				zz = zz.mul(z, z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z

				zz = zz.mul(z, z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z

				zz = zz.mul(z, z)
				zz, z = z, zz
				zz, r = zz.div(r, z, m)
				z, r = r, z
			}

			zz = zz.mul(z, powers[yi>>(_W-n)])
			zz, z = z, zz
			zz, r = zz.div(r, z, m)
			z, r = r, z

			yi <<= n
		}
	}

	return z.norm()*/
    return 1
}

//快速幂，Knuth第2卷， 4.6.3节
//Pow(x,y,m)=x^y mod m
func Pow(x,y,m int)int{
    // x**y mod 1 == 0
	if m == 1 {	return 0}
	// m == 0 || m > 1
	// x**0 == 1
	if y == 0 {	return 1}
	// y > 0
	// x**1 mod m == x mod m
	if y == 1 && m != 0 {return x % m}	
	// y > 1
// 	z:=x
	// 如果x是正整数，而y又比较大的话，我们使用4位windowed exponentiation.
	// 这样需要预先计算14个数字(x^2...x^15)，不过总的时间消耗却是原来的2/3，
	// 即使y是一个32位的指数。
	// (x^2...x^15) but then reduces the number of multiply-reduces by a
	// third. Even for a 32-bit exponent, this reduces the number of
	// operations. Uses Montgomery method for odd moduli.
	if x > 0 && y > 0 && m> 0 {
// 		if m&1 == 1 {
			return expNNMontgomery(x, y, m)
// 		}
// 		return expNNWindowed(x, y, m)
	}	
    return 0
}

/*func main(){
    //求3^1025 mod 15
    分析:
    首先，求出eular(15)=8
    则,3^8 mod 15=1(3和8互质)
    3^1025 mod 15=3^(8*128+1) mod 15=3*3^(8*128) mod 15=3*(3^8)^128 mod 15=3
  fmt.Println(Eular(15))
}*/