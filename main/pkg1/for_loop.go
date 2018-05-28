package pkg1

import (
	"fmt"
	"os"
)

/**
注释
*/
func ForLoop() {
	var s, sep string
	const a = "abcdef" // 常量
	fmt.Printf("%s\n", a)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("os.Args", os.Args[0])
	fmt.Println(sep)
}
