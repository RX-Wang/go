package pkg1

import (
	"fmt"
)

/**
常量
*/
const (
	a = iota
	b
	c
	d
	e
)

/**
iota: 是一个神奇的东西！！详见 简书
*/
func Changliang() {
	fmt.Println(a, b, c, d, e)
}
