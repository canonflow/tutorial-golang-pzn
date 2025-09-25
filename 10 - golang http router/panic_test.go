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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic: ", error)
	}
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Upps...")
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic: Upps...", string(body))
	/*
		=== RUN   TestPanicHandler
		--- PASS: TestPanicHandler (0.00s)
		PASS
	*/
}
