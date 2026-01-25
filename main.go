package main

import (
	"net/http"
	"store/controllers"
	"store/models"
	"store/repositories"
	"store/routers"
	"store/services"
)

func main() {
	categories := []models.Category{}

	categoryRepository := repositories.NewCategoryRepository(categories)
	categoryService := services.NewCategoryService(categoryRepository)
	categorieController := controllers.NewCategoryController(categoryService)
	categoryRouter := routers.NewCategoryRouter(categorieController)
	http.Handle("/api/categories/", categoryRouter.Setup())
	http.Handle("/api/categories", categoryRouter.Setup())

	if err := http.ListenAndServe(":8085", nil); err != nil {
		panic(err)
	}
}
