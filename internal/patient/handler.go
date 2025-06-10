package patient

import (
	"net/http"
	"strconv"
	"hospital-management/internal/models"
	"hospital-management/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreatePatient(c *gin.Context) {
	var req models.CreatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	userInterface := c.MustGet("user")
	user := userInterface.(map[string]interface{})
	createdBy := uint(user["id"].(float64))

	patient, err := h.service.CreatePatient(req, createdBy)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create patient", err)
		return
	}

	utils.SuccessResponse(c, "Patient created successfully", patient)
}

func (h *Handler) GetPatients(c *gin.Context) {
	patients, err := h.service.GetAllPatients()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get patients", err)
		return
	}

	utils.SuccessResponse(c, "Patients retrieved successfully", patients)
}

func (h *Handler) GetPatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	patient, err := h.service.GetPatientByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Patient not found", err)
		return
	}

	utils.SuccessResponse(c, "Patient retrieved successfully", patient)
}

func (h *Handler) UpdatePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	var req models.UpdatePatientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	patient, err := h.service.UpdatePatient(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update patient", err)
		return
	}

	utils.SuccessResponse(c, "Patient updated successfully", patient)
}

func (h *Handler) DeletePatient(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	if err := h.service.DeletePatient(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete patient", err)
		return
	}

	utils.SuccessResponse(c, "Patient deleted successfully", nil)
}

func (h *Handler) UpdateMedicalInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid patient ID", err)
		return
	}

	var req models.UpdateMedicalInfoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	patient, err := h.service.UpdateMedicalInfo(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update medical info", err)
		return
	}

	utils.SuccessResponse(c, "Medical information updated successfully", patient)
}