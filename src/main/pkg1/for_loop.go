package pkg1

import (
	// "fmt"
	// "os"
)

/**
for 循环
*/
/* func ForLoop() {
	var s, sep string
	const a = "abcdef" // 常量
	fmt.Printf("%s\n", a)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("os.Args", os.Args[0])
	fmt.Println(sep)
} */

func getG(_len int) (gs string) {
	return gs
}

func init() {
	/* // 1: 有条件的For 循环
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}
	// 2: goto 方式实现的For循环
	i := 0
	START:
		fmt.Printf("The counter is at %d\n", i)
		i++
		if i < 15 {
			goto START
		}

	// 3、只有 判断条件，没有初始值 的For循环
	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}"	*/
	// 4、什么都没有的For 循环（可谓：死循环）
	/* ii := 1
	for {
		if ii++; ii < 10 {
			fmt.Println(ii)
		} else {
			return
		}
	} */
	// 4、 for-range
	/* 	str := "Go is a beautiful language!"
	   	fmt.Printf("The length of str is: %d\n", len(str))
	   	for pos, char := range str {
	   		fmt.Printf("Character on position %d is: %c \n", pos, char)
	   	}
	   	fmt.Println()
	   	str2 := "Chinese: 汉语"
	   	fmt.Printf("The length of str2 is: %d\n", len(str2))
	   	for pos, char := range str2 {
	   			fmt.Printf("character %c starts at byte position %d\n", char, pos)
	   	}
	   	fmt.Println()
	   	fmt.Println("index int(rune) rune    char bytes")
	   	for index, rune := range str2 {
	   			fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
			 } */

	// rune对应的是每个索引的值的拷贝，所以修改rune时，原值是不会改变的
	/* iii := []int{1, 2, 3, 4, 5}
	for pos, rune := range iii {
		rune++
		fmt.Printf("position: %d   char: %d\n", pos, rune)
	}
	fmt.Println(iii) */

	/* for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	} */

	// 因为无变更条件，导致逻辑条件一直满足--> 死循环
	/* for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	} */

	/* for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	} */

	/**
			LABEL1  
	*/
	/* LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	} */
}
