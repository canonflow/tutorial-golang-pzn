package _9___golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-CANONFLOW-NAME"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-CANONFLOW-NAME")
	if err != nil {
		fmt.Fprintln(writer, "No cookie found")
	} else {
		fmt.Fprintf(writer, "Hello, %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}

}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/set-cookie?name=Nathan%20Garzya", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie Name: %s, Value: %s\n", cookie.Name, cookie.Value)
		// Cookie Name: X-CANONFLOW-NAME, Value: Nathan Garzya
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/get-cookie", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-CANONFLOW-NAME"
	cookie.Value = "Nathan Garzya"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()
	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body)) // Hello, Nathan Garzya
}
