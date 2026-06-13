package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"procurement-system/internal/logger"
	"procurement-system/internal/service"
)

type ReportHandler struct {
	service *service.ReportService
	log     logger.Logger
}

func NewReportHandler(service *service.ReportService, log logger.Logger) *ReportHandler {
	return &ReportHandler{service: service, log: log}
}

func (h *ReportHandler) Routes(r chi.Router) {
	r.Get("/inventory", h.InventoryReport)
	r.Get("/inventory/department/{department_id}", h.InventoryByDepartment)
}

func (h *ReportHandler) InventoryReport(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetInventoryReport(0)
	if err != nil {
		h.log.Error("Failed to generate inventory report", "error", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondSuccess(w, items)
}

func (h *ReportHandler) InventoryByDepartment(w http.ResponseWriter, r *http.Request) {
	deptID, err := strconv.Atoi(chi.URLParam(r, "department_id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid department ID")
		return
	}

	items, err := h.service.GetInventoryReport(deptID)
	if err != nil {
		h.log.Error("Failed to generate department report", "error", err)
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondSuccess(w, items)
}