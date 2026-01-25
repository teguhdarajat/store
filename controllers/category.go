package controllers

import (
	"encoding/json"
	"net/http"
	"store/params"
	"store/services"
)

type CategoryController interface {
	GetCategories(w http.ResponseWriter, r *http.Request)
	GetCategoryByID(w http.ResponseWriter, r *http.Request, id uint)
	InsertCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request, id uint)
	DeleteCategory(w http.ResponseWriter, r *http.Request, id uint)
}

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(service services.CategoryService) CategoryController {
	return &categoryController{
		service: service,
	}
}

func (cc *categoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories := cc.service.GetCategories()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func (cc *categoryController) GetCategoryByID(w http.ResponseWriter, r *http.Request, id uint) {
	category, err := cc.service.GetCategoryByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (cc *categoryController) InsertCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory params.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	category := cc.service.InsertCategory(newCategory)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (cc *categoryController) UpdateCategory(w http.ResponseWriter, r *http.Request, id uint) {
	var updateCategory params.Category
	err := json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	category, err := cc.service.UpdateCategory(id, updateCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}

func (cc *categoryController) DeleteCategory(w http.ResponseWriter, r *http.Request, id uint) {
	err := cc.service.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "delete succeed",
	})
}
