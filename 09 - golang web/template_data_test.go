package _9___golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Tempalte Data Struct",
		"Name":  "Canonflow",
		"Address": map[string]interface{}{
			"Street": "Jalani dulu saja",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
	/*
		<!doctype html>
			<html lang="en">
			<head>
			    <meta charset="UTF-8">
			    <meta name="viewport"
			          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
			    <meta http-equiv="X-UA-Compatible" content="ie=edge">
			    <title>Tempalte Data Struct</title>
			</head>
			<body>
			    <h1>Hello, Canonflow</h1>
			    <h2>Alamat: Jalani dulu saja</h2>
			</body>
			</html>
	*/
}

type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Tempalte Data Struct",
		Name:  "Canonflow",
		Address: Address{
			Street: "Jalani dulu saja",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Body)
	fmt.Println(string(body))
	/*
		<!doctype html>
		<html lang="en">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport"
		          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
		    <meta http-equiv="X-UA-Compatible" content="ie=edge">
		    <title>Tempalte Data Struct</title>
		</head>
		<body>
		    <h1>Hello, Canonflow</h1>
		    <h2>Alamat: Jalani dulu saja</h2>
		</body>
		</html>
	*/
}
