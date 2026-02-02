package main

import (
	"log"
	"my-first-go-project/config"
	_ "my-first-go-project/docs"
	"my-first-go-project/internal/db"
	"my-first-go-project/internal/users"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db.Connect()
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/get", users.Get)
		usersGroup.POST("/create", users.Create)
		usersGroup.PUT("/update", users.Update)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}

}
