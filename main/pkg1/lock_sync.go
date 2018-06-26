package pkg1

/**
锁 和 sync 包
*/

import "sync"
// import goex "codesite.ext/author/goExample/goex"

/**
Info 是一个共享的变量，可能被多线程操作，所以需要将其锁起
*/
type Info struct {
	mu  sync.Mutex
	Str string
}

func Update(info *Info) {
	info.mu.Lock()     // 锁住 Info 变量
	info.Str = "Diana" // new value
	info.mu.Unlock()   // 解锁
}
