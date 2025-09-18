package routes

import (
	"net/http"
	"petopia-be/controller"
	"petopia-be/service"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

// API route constants
const (
	// Product API paths
	ProductsPath       = "/api/products"
	ProductByIDPath    = "/api/products/{id:[0-9a-fA-F]+}"
	ProductByProductID = "/api/products/product/{product_id:[0-9]+}"
	AvailableProducts  = "/api/products/available"
	SearchProducts     = "/api/products/search"
	
	// Health check path
	HealthPath         = "/api/health"
	
	// Swagger path
	SwaggerPath        = "/swagger/"
)

// SetupV2 configures all routes for the application using clean architecture
func SetupV2(router *mux.Router, gormDB *gorm.DB) {
	// Create service container which will initialize all dependencies
	serviceContainer := service.NewServiceContainer()
	
	// Get the product service from the container
	productService := serviceContainer.ProductService
	
	// Initialize controllers
	productController := controller.NewProductController(productService)
	
	// Health check
	router.HandleFunc(HealthPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	}).Methods("GET")

	// Product routes with controller methods
	router.HandleFunc(ProductsPath, productController.CreateProduct).Methods("POST")
	router.HandleFunc(ProductsPath, productController.ListProducts).Methods("GET")
	router.HandleFunc(AvailableProducts, productController.GetAvailableProducts).Methods("GET")
	router.HandleFunc(SearchProducts, productController.SearchProducts).Methods("GET")
	router.HandleFunc(ProductByIDPath, productController.GetProductByID).Methods("GET")
	router.HandleFunc(ProductByProductID, productController.GetProductByProductID).Methods("GET")
	router.HandleFunc(ProductByIDPath, productController.UpdateProduct).Methods("PUT")
	router.HandleFunc(ProductByIDPath, productController.DeleteProduct).Methods("DELETE")

	// Swagger
	router.PathPrefix(SwaggerPath).Handler(httpSwagger.WrapHandler)
}
