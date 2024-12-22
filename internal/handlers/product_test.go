package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"fluffy-shop-api/internal/models"
	"fluffy-shop-api/internal/services"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	mockDB := services.NewMockDatabase()
	mockDB.Products = map[string]models.Product{
		"1": {ID: "1", Name: "Smartphone", Description: "Latest model smartphone", Price: 799.99, Stock: 10},
		"2": {ID: "2", Name: "Laptop", Description: "High-performance laptop", Price: 1200.00, Stock: 5},
	}

	handler := NewProductHandler(mockDB)

	req, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.GetProducts(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var products map[string]models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &products)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(products))
}

func TestCreateProduct(t *testing.T) {
	mockDB := services.NewMockDatabase()
	handler := NewProductHandler(mockDB)

	product := models.Product{
		ID:          "3",
		Name:        "Tablet",
		Description: "High-resolution tablet",
		Price:       499.99,
		Stock:       20,
	}

	body, _ := json.Marshal(product)
	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.CreateProduct(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	assert.Equal(t, product, mockDB.Products[product.ID])
}

func TestGetProduct(t *testing.T) {
	mockDB := services.NewMockDatabase()
	mockDB.Products = map[string]models.Product{
		"1": {ID: "1", Name: "Smartphone", Description: "Latest model smartphone", Price: 799.99, Stock: 10},
	}

	handler := NewProductHandler(mockDB)

	req, err := http.NewRequest("GET", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{productId}", handler.GetProduct)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var product models.Product
	err = json.Unmarshal(rr.Body.Bytes(), &product)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, mockDB.Products["1"], product)
}

func TestUpdateProduct(t *testing.T) {
	mockDB := services.NewMockDatabase()
	mockDB.Products = map[string]models.Product{
		"1": {ID: "1", Name: "Smartphone", Description: "Latest model smartphone", Price: 799.99, Stock: 10},
	}

	handler := NewProductHandler(mockDB)

	updatedProduct := models.Product{
		ID:          "1",
		Name:        "Updated Smartphone",
		Description: "Updated model smartphone",
		Price:       899.99,
		Stock:       15,
	}

	body, _ := json.Marshal(updatedProduct)
	req, err := http.NewRequest("PUT", "/products/1", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{productId}", handler.UpdateProduct)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	assert.Equal(t, updatedProduct, mockDB.Products["1"])
}

func TestDeleteProduct(t *testing.T) {
	mockDB := services.NewMockDatabase()
	mockDB.Products = map[string]models.Product{
		"1": {ID: "1", Name: "Smartphone", Description: "Latest model smartphone", Price: 799.99, Stock: 10},
	}

	handler := NewProductHandler(mockDB)

	req, err := http.NewRequest("DELETE", "/products/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/products/{productId}", handler.DeleteProduct)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	_, exists := mockDB.Products["1"]
	assert.False(t, exists)
}
