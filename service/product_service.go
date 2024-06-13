package service

import (
    "encoding/json"
    "go-project/model"
    "log"
    "net/http"
)

type ProductService struct {
    apiUrl string
    httpClient *http.Client
}

func NewProductService(apiUrl string) *ProductService {
    return &ProductService{
        apiUrl: apiUrl,
        httpClient: &http.Client{},
    }
}

func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
    resp, err := ps.httpClient.Get(ps.apiUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var products []model.Product
    if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
        return nil, err
    }
    return products, nil
}