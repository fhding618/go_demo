/**
 * 运行时环境变量
 * GOTRACEBACK: 开启程序崩溃时的系统核心转储功能
 * $ go build .
 * $ GOTRACEBACK=crash ./hello
 * (Ctrl+\)
 **/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world\n")
	})
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}
