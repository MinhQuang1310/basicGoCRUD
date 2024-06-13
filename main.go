package main

import (
	"basicGoCrud/db"
	"basicGoCrud/handlers"
	"basicGoCrud/models"

	"github.com/gin-gonic/gin"
)

func main() {
	gormDB := db.InitDB()
	sqlDB, err := gormDB.DB()

	if err != nil {
		panic("failed to get sql.DB")
	}
	defer sqlDB.Close()
	// Migrate schema
	gormDB.AutoMigrate(&models.TodoItem{})
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("/create", handlers.CreateItem(gormDB))
			items.GET("/getitem/:id", handlers.GetItem(gormDB))
			items.GET("/getall", handlers.GetItems(gormDB))
			items.PUT("/update/:id", handlers.UpdateItem(gormDB))
			items.DELETE("/delete/:id", handlers.DeleteItem(gormDB))
		}
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
