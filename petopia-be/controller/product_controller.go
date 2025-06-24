package controller

import (
    "encoding/json"
    "net/http"
    "gorm.io/gorm"
    "github.com/gorilla/mux"
    "strconv"
    "petopia-be/dto"
    "petopia-be/service"
)


var database *gorm.DB

func InitDB(db *gorm.DB) {
	database = db
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productRequestDTO dto.ProductRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&productRequestDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := service.CreateProduct(productRequestDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var productRequestDTO dto.ProductRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&productRequestDTO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product, err := service.UpdateProduct(id, productRequestDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	if err := service.DeleteProduct(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := service.ListProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
