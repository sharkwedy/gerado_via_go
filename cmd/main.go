package main

import (
    "go-project/config"
    "go-project/controller"
    "go-project/service"

    "github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()
    config.LoadConfig()

    productService := service.NewProductService(config.Config.ApiUrl)
    productController := controller.NewProductController(productService)

    e.GET("/api/products", productController.GetAllProducts)

    e.Logger.Fatal(e.Start(":8080"))
}