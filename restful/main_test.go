package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestJsonHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students", nil)

	mux := MakeWebHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(2, len(list))
	assert.Equal("jinsu", list[0].Name)
	assert.Equal("sungbin", list[1].Name)
}

func TestJsonHandler_second(t *testing.T) {
	assert := assert.New(t)

	var student Student
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/students/1", nil)

	mux := MakeWebHandler()

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("jinsu", student.Name)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/2", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err = json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("sungbin", student.Name)
}

func TestJsonHandler_third(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/students",
		strings.NewReader(
			`{"id":0, "name":"dongil", "age": 27, "score":88}`,
		),
	)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/3", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal("dongil", student.Name)
}

func TestJsonHandler_fourth(t *testing.T) {
	assert := assert.New(t)

	mux := MakeWebHandler()
	res := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students", nil)
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	var list []Student
	err := json.NewDecoder(res.Body).Decode(&list)
	assert.Nil(err)
	assert.Equal(1, len(list))
	assert.Equal("sungbin", list[0].Name)
}

func TestJsonHandler_fifth(t *testing.T) {
	assert := assert.New(t)

	var student Student
	mux := MakeWebHandler()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("PATCH", "/students/1",
		strings.NewReader(
			`{"score":100, "updatedScore":true}`,
		),
	)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)

	res = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/students/1", nil)

	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	err := json.NewDecoder(res.Body).Decode(&student)
	assert.Nil(err)
	assert.Equal(100, student.Score)
}
