package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Gak Ketemu")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Gak Ketemu", string(body))
	/*
		=== RUN   TestNotFoundHandler
		--- PASS: TestNotFoundHandler (0.00s)
		PASS
	*/
}
