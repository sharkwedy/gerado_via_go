package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "go-project/controller"
    "go-project/service"
    "github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
    apiUrl := "https://api.restful-api.dev/objects"
    productService := service.NewProductService(apiUrl)
    productController := controller.NewProductController(productService)

    req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
    rec := httptest.NewRecorder()

    e := echo.New()
    c := e.NewContext(req, rec)

    if assert.NoError(t, productController.GetAllProducts(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Contains(t, rec.Body.String(), "Test Product")
    }
}