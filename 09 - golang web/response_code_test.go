package _9___golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is required")
	} else {
		writer.WriteHeader(http.StatusOK)
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080/hi?name=Nathan", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("Body:", string(body))
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	/*
		===== CASE VALID =====
		Body: Hello Nathan
		Status Code: 200
		Status: 200 OK

		===== CASE INVALID =====
		Body: name is required
		Status Code: 400
		Status: 400 Bad Request
	*/
}
