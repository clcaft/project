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

type DepartmentHandler struct {
	service *service.DepartmentService
	log     logger.Logger
}

func NewDepartmentHandler(service *service.DepartmentService, log logger.Logger) *DepartmentHandler {
	return &DepartmentHandler{service: service, log: log}
}

func (h *DepartmentHandler) Routes(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
}

func (h *DepartmentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	filter := parseListFilter(r)
	items, total, err := h.service.GetAll(filter)
	if err != nil {
		h.log.Error("Failed to get departments", "error", err)
		respondError(w, http.StatusInternalServerError, "failed to get departments")
		return
	}

	respondJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Data:    items,
		Meta:    &model.Meta{Total: total, Page: filter.Page, PerPage: filter.PerPage},
	})
}

func (h *DepartmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	item, err := h.service.GetByID(id)
	if err != nil {
		h.log.Error("Failed to get department", "error", err)
		respondError(w, http.StatusInternalServerError, "failed to get department")
		return
	}
	if item == nil {
		respondError(w, http.StatusNotFound, "department not found")
		return
	}

	respondSuccess(w, item)
}

func (h *DepartmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item model.Department
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if err := h.service.Create(&item); err != nil {
		h.log.Error("Failed to create department", "error", err)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondCreated(w, item)
}

func (h *DepartmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	var item model.Department
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	item.ID = id

	if err := h.service.Update(&item); err != nil {
		h.log.Error("Failed to update department", "error", err)
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	respondSuccess(w, item)
}

func (h *DepartmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid ID")
		return
	}

	if err := h.service.Delete(id); err != nil {
		h.log.Error("Failed to delete department", "error", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondSuccess(w, map[string]string{"message": "deleted"})
}