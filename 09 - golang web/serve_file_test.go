package _9___golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileServer(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeEmbed(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
