package _9___golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	// ===== MANUAL =====
	// Cek terlebih dahulu apakah body-nya berupa Form
	err := request.ParseForm()
	if err != nil {
		return
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	// Langsung (otomatis diparsing)
	//firstName := request.PostFormValue("first_name")
	//lastName := request.PostFormValue("last_name")

	fmt.Fprintf(writer, "Hello, %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Nathan&last_name=Garzya")
	request := httptest.NewRequest("POST", "/form", requestBody)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body)) // Hello, Nathan Garzya
}
