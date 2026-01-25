package routers

import (
	"net/http"
	"store/controllers"
	"strconv"
	"strings"
)

type categoryRouter struct {
	controller controllers.CategoryController
}

func NewCategoryRouter(controller controllers.CategoryController) *categoryRouter {
	return &categoryRouter{
		controller: controller,
	}
}

func (cr *categoryRouter) Setup() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if path == "/api/categories" || path == "/api/categories/" {
			if r.Method == http.MethodGet {
				cr.controller.GetCategories(w, r)
				return
			}

			if r.Method == http.MethodPost {
				cr.controller.InsertCategory(w, r)
				return
			}
		}

		if strings.HasPrefix(path, "/api/categories/") {
			if r.Method == http.MethodGet {
				idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid Category ID", http.StatusBadRequest)
					return
				}

				cr.controller.GetCategoryByID(w, r, uint(id))
				return
			}

			if r.Method == http.MethodPut {
				idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid Category ID", http.StatusBadRequest)
					return
				}

				cr.controller.UpdateCategory(w, r, uint(id))
				return
			}

			if r.Method == http.MethodDelete {
				idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
				id, err := strconv.Atoi(idStr)
				if err != nil {
					http.Error(w, "Invalid Category ID", http.StatusBadRequest)
					return
				}

				cr.controller.DeleteCategory(w, r, uint(id))
				return
			}
		}

		http.NotFound(w, r)
	})
}
