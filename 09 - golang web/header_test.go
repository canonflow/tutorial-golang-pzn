package _9___golang_web

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type") // No case-sensitive
	fmt.Fprintf(writer, "Content-Type: "+contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	// Add new Header
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body)) // Content-Type: application/x-www-form-urlencoded
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "canonflow")
	fmt.Fprintf(writer, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))                        // OK
	fmt.Println(response.Header.Get("X-Powered-By")) // canonflow

}
