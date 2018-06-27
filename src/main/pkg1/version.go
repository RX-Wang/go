package pkg1

import (
	"fmt"
	"runtime"
)

/**
注释
*/
func Version() {
	fmt.Printf("go`s version:%s\n", runtime.Version())
}
