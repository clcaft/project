package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"procurement-system/internal/logger"
	"procurement-system/internal/model"
	"procurement-system/internal/service"
)

type InventoryBalanceHandler struct {
	service *service.InventoryBalanceService
	log     logger.Logger
}

func NewInventoryBalanceHandler(service *service.InventoryBalanceService, log logger.Logger) *InventoryBalanceHandler {
	return &InventoryBalanceHandler{service: service, log: log}
}

func (h *InventoryBalanceHandler) Routes(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Get("/department/{department_id}", h.GetByDepartment)
}

func (h *InventoryBalanceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	filter := parseListFilter(r)
	items, total, err := h.service.GetAll(filter)
	if err != nil {
		h.log.Error("Failed to get inventory balances", "error", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Data:    items,
		Meta:    &model.Meta{Total: total, Page: filter.Page, PerPage: filter.PerPage},
	})
}

func (h *InventoryBalanceHandler) GetByDepartment(w http.ResponseWriter, r *http.Request) {
	deptID, err := strconv.Atoi(chi.URLParam(r, "department_id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid department ID")
		return
	}

	items, err := h.service.GetByDepartment(deptID)
	if err != nil {
		h.log.Error("Failed to get department balances", "error", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondSuccess(w, items)
}