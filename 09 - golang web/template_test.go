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

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	// SIMPLE -> nama templatenya (karena kita pakai string)
	t, err := template.New("SIMPLE").Parse(templateText)

	if err != nil {
		panic(err)
	}

	// Pakai ini otomatis, tidak perlu pengecekan err != nill
	//t := template.Must(template.New("SIMPLE").Parse(templateText))

	err = t.ExecuteTemplate(writer, "SIMPLE", "Hello, I'm Canonflow!")
	if err != nil {
		panic(err)
	}
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body)) // <html><body>Hello, I&#39;m Canonflow!</body></html>
}

func SimpleFileHTML(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/simple.gohtml")

	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(writer, "simple.gohtml", "Hello, I'm Canonflow!")
	if err != nil {
		return
	}
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")

	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(writer, "simple.gohtml", "Hello, I'm Canonflow!")
	if err != nil {
		return
	}
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFS(templates, "templates/*.gohtml")

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello, I'm Canonflow!")
}

func TestSimpleFileHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleFileHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
	/*
		<!doctype html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport"
		          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
		    <meta http-equiv="X-UA-Compatible" content="ie=edge">
		    <title>Hello, I&#39;m Canonflow!</title>
		</head>
		<body>
		<h1>Hello, I&#39;m Canonflow!</h1>
		</body>
		</html>
	*/
}

func TestSimpleTemplateDirecotry(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
	/*
		<!doctype html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport"
		          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
		    <meta http-equiv="X-UA-Compatible" content="ie=edge">
		    <title>Hello, I&#39;m Canonflow!</title>
		</head>
		<body>
		<h1>Hello, I&#39;m Canonflow!</h1>
		</body>
		</html>
	*/
}

func TestSimpleTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
	/*
		<!doctype html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport"
		          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
		    <meta http-equiv="X-UA-Compatible" content="ie=edge">
		    <title>Hello, I&#39;m Canonflow!</title>
		</head>
		<body>
		<h1>Hello, I&#39;m Canonflow!</h1>
		</body>
		</html>
	*/
}
