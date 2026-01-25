package main

import (
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}
}
