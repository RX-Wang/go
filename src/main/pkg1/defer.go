package pkg1

import (
	"fmt"
	"time"
)

func init() {
	// defer01();
	//defer02();
	// defer04();
	defer08();
}

/**
 输出：
	打印--1： 1
	打印--2： 0
	打印--3： 1
	打印--4： 0
*/
func defer01() {
	i := 0
	defer fmt.Println("defer--i:", i)
	i++
	return
}
func defer02() {
	i := 0
	defer defer03(i)
	defer func() {
		fmt.Println("打印--3：", i)
	}()
	defer fmt.Println("打印--2：", i)
	i++
	fmt.Println("打印--1：", i)
}
func defer03(i int) {
	fmt.Println("打印--4：", i)
}

/**
 输出：
	打印--1： abcdef
	打印--2： abc
	打印--3： abcdef
	打印--4： abc
*/
func defer04() {
	i := "abc"
	defer defer05(i)
	defer func() {
		fmt.Println("打印--3：", i)
	}()
	defer fmt.Println("打印--2：", i)
	i += "def"
	fmt.Println("打印--1：", i)
}
func defer05(i string) {
	fmt.Println("打印--4：", i)
}


func defer06(funcName string) func(){
	start := time.Now()
	fmt.Printf("function %s enter\n",funcName)
	return func(){
		fmt.Printf("function %s exit (elapsed %s)\n",funcName,time.Since(start))
	}
}

func defer07(funcName string){
	defer defer06(funcName)()
	time.Sleep(2 * time.Second)
}
func defer08(){
	defer07("func01")
	defer07("func02")
}