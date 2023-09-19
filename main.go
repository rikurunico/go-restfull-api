package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	"github.com/rikurunico/go-restfull-api/app"
	"github.com/rikurunico/go-restfull-api/controller"
	"github.com/rikurunico/go-restfull-api/helper"
	"github.com/rikurunico/go-restfull-api/repository"
	"github.com/rikurunico/go-restfull-api/service"

)

func main() {
	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:id", categoryController.Update)
	router.DELETE("/api/categories/:id", categoryController.Delete)
	router.GET("/api/categories/:id", categoryController.FindById)
	router.GET("/api/categories", categoryController.FindAll)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	fmt.Println("Server is running at localhost:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
