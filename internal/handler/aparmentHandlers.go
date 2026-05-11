package handler

import (
	"net/http"

	"github.com/adocoder12/Costabackend/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *App) CreateApartmentHandler(c *gin.Context) {
	var req dto.CreateApartmentRequest

	// 1. Bind & Validate Request DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		app.Logger.Warn("invalid apartment request", "error", err)
		app.clientError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 2. Call Service with the DTO directly
	// The service returns the already-mapped Response DTO
	response, err := app.ApartmentService.CreateApartment(c.Request.Context(), req)

	if err != nil {
		app.serverError(c, err)
		return
	}

	// 3. Return the Response
	c.JSON(http.StatusCreated, response)
}

func (app *App) GetApartmentsHandler(c *gin.Context) {
	// 1. Call Service (which now handles mapping and IsMine logic)
	viewerID := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	apartments, err := app.ApartmentService.GetApartments(c.Request.Context(), viewerID)

	// 2. Error handling
	if err != nil {
		app.Logger.Error("failed to fetch apartments", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve apartments"})
		return
	}

	// 3. Just Return!
	// 'apartments' is already []dto.ApartmentResponse. No more mapping needed here.
	c.JSON(http.StatusOK, apartments)
}
