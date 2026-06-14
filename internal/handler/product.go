package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"procurement-system/internal/logger"
	"procurement-system/internal/model"
	"procurement-system/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
	log     logger.Logger
}

func NewProductHandler(service *service.ProductService, log logger.Logger) *ProductHandler {
	return &ProductHandler{service: service, log: log}
}

func (h *ProductHandler) Routes(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	filter := parseListFilter(r)
	items, total, err := h.service.GetAll(filter)
	if err != nil {
		h.log.Error("Failed to get products", "error", err)
		respondError(w, http.StatusInternalServerError, "failed to get products")
		return
	}

	respondJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Data:    items,
		Meta:    &model.Meta{Total: total, Page: filter.Page, PerPage: filter.PerPage},
	})
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	item, err := h.service.GetByID(id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if item == nil {
		respondError(w, http.StatusNotFound, "not found")
		return
	}
	respondSuccess(w, item)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item model.Product
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}
	if err := h.service.Create(&item); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondCreated(w, item)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var item model.Product
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, "invalid body")
		return
	}
	item.ID = id
	if err := h.service.Update(&item); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondSuccess(w, item)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondSuccess(w, map[string]string{"message": "deleted"})
}