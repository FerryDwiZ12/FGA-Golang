package main

import (
	"github.com/gin-gonic/gin"
	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/config"
	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/internal/order"
	"github.com/FerryDwiZ12/FGA-Golang/tree/master/assigments/pkg/datasource"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("local")
	if err != nil {
		panic(err)
	}

	// Create Connection Database & Migrate
	db, err := datasource.NewDatabase(cfg.Database)
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	orderRepo := order.NewRepository(db)
	orderSvc := order.NewService(orderRepo)
	orderHandler := order.NewHandler(orderSvc)

	// Route
	app.POST("/orders", orderHandler.CreateOrder)
	app.GET("/orders", orderHandler.GetOrder)
	app.PUT("/orders/:id", orderHandler.UpdateOrder)
	app.DELETE("/orders/:id", orderHandler.DeleteOrder)

	app.Run(":8080")
}
