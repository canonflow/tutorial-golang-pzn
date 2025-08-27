package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"golang-dependency-injection/app"
	"golang-dependency-injection/controller"
	"golang-dependency-injection/helper"
	"golang-dependency-injection/middleware"
	"golang-dependency-injection/repository"
	"golang-dependency-injection/service"
	"net/http"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
