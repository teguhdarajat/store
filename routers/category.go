package routers

import (
	"net/http"
	"store/handlers"
)

func RegisterCategoryRoutes(categoryController handlers.CategoryHandler) {
	http.HandleFunc("/api/categories", categoryController.HandleCategories)
	http.HandleFunc("/api/category/", categoryController.HandleCategoryByID)
}
