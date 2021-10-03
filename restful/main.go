package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

var (
	students map[int]Student
	lastId   int
)

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")

	students = make(map[int]Student)
	students[1] = Student{1, "jinsu", 27, 87}
	students[2] = Student{2, "sungbin", 20, 99}
	lastId = 2

	return mux
}

func GetStudentListHandler(writer http.ResponseWriter, request *http.Request) {
	list := make([]Student, 0)
	for _, student := range students {
		list = append(list, student)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Id < list[j].Id
	})

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(list)
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
