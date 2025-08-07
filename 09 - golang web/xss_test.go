package _9___golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>Ini adalah body</p>"), // Akan di-render sebagai HTML biasa
	})
}

func TestTemplateAutoEscape(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	recorder := httptest.NewRecorder()
	TemplateAutoEscape(recorder, request)
	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
}
