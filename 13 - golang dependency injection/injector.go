//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"golang-dependency-injection/app"
	"golang-dependency-injection/controller"
	"golang-dependency-injection/middleware"
	"golang-dependency-injection/repository"
	"golang-dependency-injection/service"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	// Kalau ada yang butuh CategoryRepository, kirim CategoryRepositoryImpl
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	// Kalau ada yang butuh CategoryService, kirim CategoryServiceImpl
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	// Kalau ada yang butuh CategoryController, kirim CategoryControllerImpl
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		NewValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
