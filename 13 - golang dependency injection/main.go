package main

import (
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"golang-dependency-injection/helper"
	"golang-dependency-injection/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":3000",
		Handler: authMiddleware,
	}
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func main() {

	//validate := validator.New()
	//db := app.NewDB()
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//
	//router := app.NewRouter(categoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)
	//
	//server := NewServer(authMiddleware)
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
