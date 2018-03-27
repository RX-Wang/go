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
	fmt.Print(a)
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
