package controller

import (
    "net/http"
    "go-project/service"
    "github.com/labstack/echo/v4"
)

type ProductController struct {
    service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
    return &ProductController{service: service}
}

func (pc *ProductController) GetAllProducts(c echo.Context) error {
    products, err := pc.service.GetAllProducts()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, products)
}