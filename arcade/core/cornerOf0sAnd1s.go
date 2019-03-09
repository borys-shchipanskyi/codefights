package core

import "fmt"

func DecToBin(n int)int{
	if n <= 1{
		return n
	}
	return 10 * DecToBin(n/2) + n % 2
}

func DecToBinMir(n int)int{
	res := 0
	 for n > 0{
		 res = (res + n % 2) * 10
		 n /= 2
	 }
	res /= 10
	return res
}

func Pow2(n int) int{
	res := 1
	for i := 1; i <= n; i++{
		res *= 2
	}
	return res
}

func BinToDec(n, i int) int{
	res := 0
	for n > 0{
		if n % 10 == 1{
			res += Pow2(i)
		}
		n /= 10
		i++
	}
	return res

}

func ArrayPacking(a []int) int {
	res := 0
	for i := 0; i <len(a); i++{
		binVal := DecToBin(a[i])
		res += BinToDec(binVal, i * 8)
	}
	return res
}

func get1Num(n int) int{
	count:= 0
	for n > 0 {
		if n % 2 == 1{
			count++
		}
		n /= 2
	}

	return count
}
func RangeBitCount(a int, b int) int {
	res := 0
	for i := a; i <= b; i++{

		res +=  get1Num(i)
		fmt.Println(i, get1Num(i))
	}
	return res
}

func CountSumOfTwoRepresentations2(n int, l int, r int) int {
	count := 0
	for i := l; i <= r && n >i; i++{
		res := n - i
		if res >= l && res <= r && res >= i{
			count++
		}
	}
	return count
}

func AdditionWithoutCarrying(param1 int, param2 int) int {
	fmt.Println(param1, param2)
	res := 0
	mv := 1
	for param1 > 0 && param2 > 0{
		res = res + (param1 % 10 + param2 % 10)%10 * mv
		fmt.Println(res, param1 % 10, param2 % 10, (param1 % 10 + param2 % 10)%10)
		param1 /= 10
		param2 /= 10
		mv *= 10
	}
	if param1 == 0 && param2 != 0{
		res = param2 * mv + res
	}
	if param2 == 0 && param1 != 0{
		res = param1 * mv + res
	}
	fmt.Println(res, param1, param2)
	return res
}