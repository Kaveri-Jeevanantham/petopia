package controller

import (
	"encoding/json"
	"net/http"
	"petopia-be/dto"
	"petopia-be/service"
	"strconv"

	"github.com/gorilla/mux"
)

// Constants for controller responses
const (
	ContentTypeHeaderKey = "Content-Type"
	ContentTypeJSONValue = "application/json"
	ErrProductNotFound   = "product not found"
	ErrInvalidProductID  = "invalid product ID"
)

// ProductController handles the HTTP requests for products
type ProductController struct {
	ProductService service.ProductService
}

// NewProductController creates a new product controller
func NewProductController(ps service.ProductService) *ProductController {
	return &ProductController{
		ProductService: ps,
	}
}

// CreateProduct handles product creation
func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var requestDTO dto.ProductRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call service layer
	createdProduct, err := c.ProductService.CreateProduct(r.Context(), requestDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

// GetProductByID retrieves a product by its ID
func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	product, err := c.ProductService.GetProductByID(r.Context(), id)
	if err != nil {
		if err.Error() == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err.Error() == ErrInvalidProductID {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	json.NewEncoder(w).Encode(product)
}

// GetProductByProductID retrieves a product by its product_id
func (c *ProductController) GetProductByProductID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr := vars["product_id"]

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := c.ProductService.GetProductByProductID(r.Context(), productID)
	if err != nil {
		if err.Error() == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	json.NewEncoder(w).Encode(product)
}

// ListProducts returns a list of products with pagination
func (c *ProductController) ListProducts(w http.ResponseWriter, r *http.Request) {
	page := int64(1)
	limit := int64(10)

	queryParams := r.URL.Query()

	if pageStr := queryParams.Get("page"); pageStr != "" {
		if pageNum, err := strconv.ParseInt(pageStr, 10, 64); err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	if limitStr := queryParams.Get("limit"); limitStr != "" {
		if limitNum, err := strconv.ParseInt(limitStr, 10, 64); err == nil && limitNum > 0 {
			limit = limitNum
		}
	}

	category := queryParams.Get("category")
	if category != "" {
		// If a category is specified, get products by category
		products, err := c.ProductService.GetProductsByCategory(r.Context(), category)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
		json.NewEncoder(w).Encode(products)
		return
	}

	// Get all products with pagination
	result, err := c.ProductService.GetAllProducts(r.Context(), page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	json.NewEncoder(w).Encode(result)
}

// GetAvailableProducts returns only available products
func (c *ProductController) GetAvailableProducts(w http.ResponseWriter, r *http.Request) {
	page := int64(1)
	limit := int64(10)

	queryParams := r.URL.Query()

	if pageStr := queryParams.Get("page"); pageStr != "" {
		if pageNum, err := strconv.ParseInt(pageStr, 10, 64); err == nil && pageNum > 0 {
			page = pageNum
		}
	}

	if limitStr := queryParams.Get("limit"); limitStr != "" {
		if limitNum, err := strconv.ParseInt(limitStr, 10, 64); err == nil && limitNum > 0 {
			limit = limitNum
		}
	}

	result, err := c.ProductService.GetAvailableProducts(r.Context(), page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	json.NewEncoder(w).Encode(result)
}

// SearchProducts searches for products by term
func (c *ProductController) SearchProducts(w http.ResponseWriter, r *http.Request) {
	term := r.URL.Query().Get("q")
	if term == "" {
		http.Error(w, "Search term is required", http.StatusBadRequest)
		return
	}

	products, err := c.ProductService.SearchProducts(r.Context(), term)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	json.NewEncoder(w).Encode(products)
}

// UpdateProduct updates an existing product
func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var requestDTO dto.ProductRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := c.ProductService.UpdateProduct(r.Context(), id, requestDTO)
	if err != nil {
		if err.Error() == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err.Error() == ErrInvalidProductID {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the updated product
	updatedProduct, err := c.ProductService.GetProductByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set(ContentTypeHeaderKey, ContentTypeJSONValue)
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProduct deletes a product
func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := c.ProductService.DeleteProduct(r.Context(), id)
	if err != nil {
		if err.Error() == ErrProductNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err.Error() == ErrInvalidProductID {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
