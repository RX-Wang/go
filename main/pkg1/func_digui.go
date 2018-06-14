package pkg1

/**
递归 的demo
*/

import (
	// "strings"
	"fmt"
	// "unicode"
	"time"
)

const LIM = 41

var fibs [LIM + 1]uint64

func init() {
	// 1.1、斐波那契 数列---递归形式
	/* start := time.Now()
	ind, fib := fibonacci(41)
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("the number is : %d，fib is : %d\n", ind, fib)
	fmt.Printf("delta is : %s\n", delta) */

	// 1.2、斐波那契 数列  -- 闭包方式
	// fmt.Printf("fibonacci2 : %d\n", fibonacci2(10))


	// 1.3、基于内存缓存实现的菲波那切数列
	/* var result uint64
	start1 := time.Now()
	for i := 0; i < LIM+1; i++ {
		result = fibonacci3(i)
	}
	end1 := time.Now()
	delta1 := end1.Sub(start1)
	fmt.Printf("fibonacci(%d) is: %d\n", LIM, result)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta1) */


	// 2、递归打印  10-1
	// printNum(10)

	// 3、阶乘
	// ind, res := jc(4)
	// fmt.Printf("the number: %d  jc is : %d\n", ind, res)

	// 4、回调函数的传递
	// st := "hello world !"
	// fmt.Println(strings.IndexFunc(st, unicode.IsSpace))

	// 5.1、闭包--自执行函数
	/* sum := 0
	func() {
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}()
	fmt.Println(sum) */

	// 5.2、 匿名函数赋值给变量
	/* for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	} */

	// 5.3、 匿名函数与defer + Println 与 defer 的区别
	/* func() (res int){
		defer func() { fmt.Println("res is : ", res + 1) }()
		return 1
	}()

	func() {
		i := 1
		defer func() {
			fmt.Println("defer non-declaration func Println ", i)
		}()
		i++
		fmt.Println("non-declaration func Println ", i)
	}()

	func() {
		i := 1
		defer fmt.Println("defer Println ", i)
		i++
		fmt.Println("normal Println ", i)
	}() */

	// 5.4、匿名函数 作为返回值使用
	// fmt.Println("匿名函数作为返回值：", noNameFuncReturn(2)(3))

	// 6.1、函数运行时间
	// funcRunedTime()
}

// 1.1、斐波那契 数列  -- 递归方式
func fibonacci(n int) (ind int, res int) {
	ind = n
	if n <= 1 {
		res = 1
	} else {
		_, fib1 := fibonacci(n - 1)
		_, fib2 := fibonacci(n - 2)
		res = fib1 + fib2
	}
	return
}

// 1.2、斐波那契 数列  -- 闭包方式
func fibonacci2(n int) int {
	if n == 0 {
		return 0
	} else {
		var res int
		ff := func() func() int {
			back1, back2 := 1, 1 // 预先定义好前两个值
			return func() int {
				//记录（back1）的值
				temp := back1
				// 重新赋值(这个就是核心代码)
				back1, back2 = back2, (back1 + back2)
				//返回temp
				return temp
			}
		}()
		for i := 1; i <= n+1; i++ {
			if i == n+1 {
				res = ff()
			} else {
				ff()
			}
		}
		return res
	}
}

// 1.3、基于内存缓存实现的菲波那切数列
func fibonacci3(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci3(n-1) + fibonacci3(n-2)
	}
	fibs[n] = res
	return
}

// 2、递归打印  10-1

func printNum(num int) {
	if num == 1 {
		fmt.Println(num)
	} else {
		fmt.Println(num)
		printNum(num - 1)
	}
}

// 3、阶乘

func jc(num int) (ind int, res int) {
	ind = num
	if num == 0 {
		res = 1
	} else {
		_, res1 := jc(num - 1)
		res = num * res1
	}
	return
}

// 5.4、匿名函数 作为返回值使用

func noNameFuncReturn(n int) func(m int) int {
	return func(m int) int {
		return n + m
	}
}

// 6.1、函数执行时间
func funcRunedTime() {
	start := time.Now()
	fmt.Println("菲波那切数列：", fibonacci2(15))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
