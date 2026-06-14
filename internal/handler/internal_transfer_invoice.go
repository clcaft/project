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

type InternalTransferInvoiceHandler struct {
	service *service.InternalTransferInvoiceService
	log     logger.Logger
}

func NewInternalTransferInvoiceHandler(service *service.InternalTransferInvoiceService, log logger.Logger) *InternalTransferInvoiceHandler {
	return &InternalTransferInvoiceHandler{service: service, log: log}
}

func (h *InternalTransferInvoiceHandler) Routes(r chi.Router) {
	r.Get("/", h.GetAll)
	r.Get("/{id}", h.GetByID)
	r.Post("/", h.Create)
	r.Put("/{id}", h.Update)
	r.Delete("/{id}", h.Delete)
	r.Post("/{id}/confirm", h.Confirm)
	r.Post("/{id}/ship", h.Ship)
	r.Post("/{id}/receive", h.Receive)
	r.Post("/{id}/cancel", h.Cancel)
}

func (h *InternalTransferInvoiceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	filter := parseListFilter(r)
	items, total, err := h.service.GetAll(filter)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Data:    items,
		Meta:    &model.Meta{Total: total, Page: filter.Page, PerPage: filter.PerPage},
	})
}

func (h *InternalTransferInvoiceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
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

func (h *InternalTransferInvoiceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var item model.InternalTransferInvoice
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

func (h *InternalTransferInvoiceHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var item model.InternalTransferInvoice
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

func (h *InternalTransferInvoiceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Delete(id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondSuccess(w, map[string]string{"message": "deleted"})
}

func (h *InternalTransferInvoiceHandler) Confirm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Confirm(id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondSuccess(w, map[string]string{"status": "confirmed"})
}

func (h *InternalTransferInvoiceHandler) Ship(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Ship(id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondSuccess(w, map[string]string{"status": "shipped"})
}

func (h *InternalTransferInvoiceHandler) Receive(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Receive(id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondSuccess(w, map[string]string{"status": "received"})
}

func (h *InternalTransferInvoiceHandler) Cancel(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.service.Cancel(id); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondSuccess(w, map[string]string{"status": "cancelled"})
}