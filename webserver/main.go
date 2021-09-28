package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func studentHandler(writer http.ResponseWriter, req *http.Request){
	var student = Student{Name: "jinsu", Age: 27, Score: 99}
	data, _ := json.Marshal(student)
	writer.Header().Add("content-type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, string(data))
}

func barHandler(writer http.ResponseWriter, req *http.Request) {
	values := req.URL.Query()  // 쿼리 인수
	name := values.Get("name") // 특정 키 값 확인
	if name == "" {
		name = "World"
	}
	id, _ := strconv.Atoi(values.Get("id")) // id 값을 int 형으로 변환
	fmt.Fprintf(writer, "Hello %s! id: %d", name, id)
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()

	// 웹 핸들러 등록
	mux.HandleFunc("/", func(writer http.ResponseWriter, req *http.Request) {
		fmt.Fprint(writer, "Hello World!")
	})
	mux.HandleFunc("/bar", barHandler)
	mux.HandleFunc("/student", studentHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	return mux
}

func main() {

	// 웹 서버 시작
	http.ListenAndServe(":3000", MakeWebHandler())
}
