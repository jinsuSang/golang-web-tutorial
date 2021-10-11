package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"sort"
	"strconv"
)

var rd *render.Render

type Todo struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name"`
	Completed bool   `json:"completed,omitempty"`
}

type Success struct {
	Success bool `json:"success"`
}

var todoMap map[int]Todo
var lastID int = 0

func MakeWebHandler() http.Handler {
	todoMap = make(map[int]Todo)

	mux := mux.NewRouter()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/todos", GetTodoListHandler).Methods("GET")
	mux.HandleFunc("/todos", PostTodoListHandler).Methods("POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", RemoveTodoHandler).Methods("DELETE")
	mux.HandleFunc("/todos/{id:[0-9]+}", UpdateTodoHandler).Methods("PUT")

	return mux
}

func GetTodoListHandler(writer http.ResponseWriter, request *http.Request) {
	list := make([]Todo, 0)
	for _, todo := range todoMap {
		list = append(list, todo)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID > list[i].ID
	})
	rd.JSON(writer, http.StatusOK, list)
}

func PostTodoListHandler(writer http.ResponseWriter, request *http.Request) {
	var todo Todo
	err := json.NewDecoder(request.Body).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	lastID++
	todo.ID = lastID
	todoMap[lastID] = todo
	rd.JSON(writer, http.StatusCreated, todo)
}

func RemoveTodoHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id)
		rd.JSON(writer, http.StatusOK, Success{true})
	} else {
		rd.JSON(writer, http.StatusNotFound, Success{false})
	}
}

func UpdateTodoHandler(writer http.ResponseWriter, request *http.Request) {
	var newTodo Todo
	err := json.NewDecoder(request.Body).Decode(&newTodo)
	if err != nil {
		log.Fatal(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	if todo, ok := todoMap[id]; ok {
		todo.Name = newTodo.Name
		todo.Completed = newTodo.Completed
		rd.JSON(writer, http.StatusOK, Success{true})
	} else {
		rd.JSON(writer, http.StatusBadRequest, Success{false})
	}
}

func main() {
	rd = render.New()
	m := MakeWebHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
