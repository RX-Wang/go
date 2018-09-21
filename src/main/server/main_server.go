package main

import (
	"fmt"
	"log"
	"net/http"
)
// 原生 服务
func helloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+ req.URL.Path)
}

func main() {
	fmt.Println("go server 8081")
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe("localhost:8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/sirupsen/logrus"
// )

// func handleWebhook(c *gin.Context) {
// }
// func main() {
// 	var err error
// 	r := gin.Default()
// 	r.GET("/webhook", handleWebhook)
// 	if err = r.Run("127.0.0.1:8081"); err != nil {
// 		logrus.WithError(err).Fatal("Couldn't start server")
// 	}
// }
