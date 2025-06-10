package user

import (
	"net/http"
	"hospital-management/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// @Summary Get user profile
// @Description Get current user profile
// @Tags user
// @Security Bearer
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /profile [get]
func (h *Handler) GetProfile(c *gin.Context) {
	userInterface := c.MustGet("user")
	userMap := userInterface.(map[string]interface{})
	userID := uint(userMap["id"].(float64))

	user, err := h.service.GetByID(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found", err)
		return
	}

	utils.SuccessResponse(c, "Profile retrieved successfully", user)
}

// @Summary Update user profile
// @Description Update current user profile
// @Tags user
// @Security Bearer
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /profile [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	userInterface := c.MustGet("user")
	userMap := userInterface.(map[string]interface{})
	userID := uint(userMap["id"].(float64))

	user, err := h.service.GetByID(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found", err)
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	// Only update allowed fields
	if firstName, ok := updateData["first_name"].(string); ok {
		user.FirstName = firstName
	}
	if lastName, ok := updateData["last_name"].(string); ok {
		user.LastName = lastName
	}
	if phone, ok := updateData["phone"].(string); ok {
		user.Phone = phone
	}
	if email, ok := updateData["email"].(string); ok {
		user.Email = email
	}

	if err := h.service.Update(user); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update profile", err)
		return
	}

	utils.SuccessResponse(c, "Profile updated successfully", user)
}
