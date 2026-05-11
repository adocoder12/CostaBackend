package handler

import (
	"net/http"

	"github.com/adocoder12/Costabackend/internal/dto"
	"github.com/gin-gonic/gin"
)

func (app *App) CreateApartmentHandler(c *gin.Context) {
	var req dto.CreateApartmentRequest

	// 1. Bind & Validate Request DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. Use the Mapper to get a Model
	apartmentModel := req.ToModel()

	// 3. Service Logic
	if err := app.ApartmentService.CreateApartment(c.Request.Context(), apartmentModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save"})
		return
	}

	// 4. Return the Response DTO
	c.JSON(http.StatusCreated, dto.FromModel(apartmentModel))
}
