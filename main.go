package main

import (
	"fmt"
	"log"
	"net/http"
	"store/config"
	"store/database"
	"store/handlers"
	"store/repositories"
	"store/routers"
	"store/services"
)

func main() {
	config := config.LoadConfig()
	db, err := database.InitDB(config.Database)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	routers.RegisterCategoryRoutes(categoryHandler)

	fmt.Println(config.Database)

	if err := http.ListenAndServe("8085", nil); err != nil {
		panic(err)
	}
}
