package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 웹 핸들러 등록
	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request){
		fmt.Fprint(writer, "Hello World")
	})
	// 웹 서버 시작
	http.ListenAndServe(":3000", nil)
}
