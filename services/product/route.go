package product

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"vitshop.vimfn.in/types"
	"vitshop.vimfn.in/utils"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/listing", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/listing/{productID}", h.handleGetProductById).Methods(http.MethodGet)

	router.HandleFunc("/listing", h.handleCreateProduct).Methods(http.MethodPost)

	router.HandleFunc("/listing/{productID}", h.handleUpdateProduct).Methods(http.MethodPut)
	router.HandleFunc("/listing/{productID}", h.handleDeleteProduct).Methods(http.MethodDelete)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data": products,
	})
}

func (h *Handler) handleGetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}
	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}
	product, err := h.store.GetProductByID(productID)
	if err != nil {
		// check if it's a "not found" error
		if strings.Contains(err.Error(), "not found") {
			utils.WriteError(w, http.StatusNotFound, err)
		} else {
			utils.WriteError(w, http.StatusInternalServerError, err)
		}
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"data": product,
	})
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProdPayload types.CreateProductPayload

	// fmt.Println(r)

	if err := utils.ParseJSON(r, &createProdPayload); err != nil {
		fmt.Println("fked in parsing")
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(createProdPayload); err != nil {
		fmt.Println("fked in validation")
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	newProd, err := h.store.CreateProduct(createProdPayload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated,
		map[string]interface{}{
			"data": newProd,
		})
}

func (h *Handler) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	if err := h.store.DeleteProductByID(productID); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	str, ok := vars["productID"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing product ID"))
		return
	}

	productID, err := strconv.Atoi(str)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid product ID"))
		return
	}

	var productPayload types.UpdateProductPayload
	if err := utils.ParseJSON(r, &productPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(productPayload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	updatedProduct, err := h.store.UpdateProduct(productID, productPayload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated,
		map[string]interface{}{
			"data": updatedProduct,
		})
}
