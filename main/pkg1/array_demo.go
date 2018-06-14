/**
数组和切片
*/

package pkg1

import (
	"fmt"
)

func init() {
	/* a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	} */
	var arr1 [5]int
	var arr2 = new([5]int) // new 出来的是一个指针
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
}
