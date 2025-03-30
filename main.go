package main

import (
	"afryn123/golang-restful-api/app"
	"afryn123/golang-restful-api/controller"
	"afryn123/golang-restful-api/exception"
	"afryn123/golang-restful-api/helper"
	"afryn123/golang-restful-api/middleware"
	"afryn123/golang-restful-api/repository"
	"afryn123/golang-restful-api/service"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	var port string = "3000"
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryServiceImpl(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.POST("/api/categories", categoryController.Create)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:" + port,
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server listening on " + port)
	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
