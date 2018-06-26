/**
协程
*/

package pkg1

import (
	"fmt"
	"time"
)

type RemoteResult struct {
	Url    string
	Result string
}

func remoteGet(requestUrl string, resultChan chan RemoteResult) {
	resultChan <- RemoteResult{Url: requestUrl, Result: requestUrl}
}
func multiGet(urls []string) []RemoteResult {
	fmt.Println(time.Now())
	resultChan := make(chan RemoteResult, len(urls))
	defer close(resultChan) // 在所有代码执行完毕以后 延迟关闭协程
	var result []RemoteResult
	//fmt.Println(result)
	for _, url := range urls {
		go remoteGet(url, resultChan) // 开启协程
	}
	for i := 0; i < len(urls); i++ {
		res := <-resultChan
		result = append(result, res)
	}
	fmt.Println(time.Now())
	return result
}

func init() {
	/* urls := []string{
		"http://127.0.0.1/test.php?i=13\n",
		"http://127.0.0.1/test.php?i=14\n",
		"http://127.0.0.1/test.php?i=15\n",
		"http://127.0.0.1/test.php?i=16\n",
		"http://127.0.0.1/test.php?i=17\n",
		"http://127.0.0.1/test.php?i=18\n",
		"http://127.0.0.1/test.php?i=19\n",
		"http://127.0.0.1/test.php?i=20\n"}
	content := multiGet(urls)
	fmt.Println(content) */
}
