package _9___golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var templates_ embed.FS

var myTemplates = template.Must(template.ParseFS(templates_, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello template caching")
}

func TestTemplateCaching(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()
	TemplateCaching(recorder, request)
	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
}
