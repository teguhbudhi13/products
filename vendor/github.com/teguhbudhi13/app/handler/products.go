package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/teguhbudhi13/products/app/model"
)

// GetAllproducts get all product data
func GetAllproducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Find(&products)
	respondJSON(w, http.StatusOK, products)
}

// Createproduct add new product
func Createproduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

// Getproduct get specific project
func Getproduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	product := getproductByID(db, id, w, r)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	respondJSON(w, http.StatusOK, product)
}

// Updateproduct update spesific product data
func Updateproduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	product := getproductByID(db, id, w, r)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, product)
}

// Deleteproduct delete specific product
func Deleteproduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])
	product := getproductByID(db, id, w, r)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := db.Delete(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getproductByID gets a product instance if exists
func getproductByID(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *model.Product {
	product := model.Product{}
	if err := db.First(&product, id).Error; err != nil {
		return nil
	}
	return &product
}
