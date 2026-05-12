package handler

import (
	"errors"
	"net/http"

	"github.com/adocoder12/Costabackend/internal/dto"
	"github.com/adocoder12/Costabackend/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *App) GetApartmentsHandler(c *gin.Context) {
	// TODO: replace with c.GetString("userID") once auth middleware is wired
	viewerID := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	apartments, err := app.ApartmentService.GetApartments(c.Request.Context(), viewerID)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, apartments)
}

func (app *App) GetApartmentByIDHandler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		app.clientError(c, http.StatusBadRequest, "invalid apartment id format")
		return
	}

	// TODO: replace with c.GetString("userID") once auth middleware is wired
	viewerID := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	apartment, err := app.ApartmentService.GetApartmentByID(c.Request.Context(), id, viewerID)
	if err != nil {
		// FIX: use errors.Is with sentinel — never string match on error messages
		if errors.Is(err, repository.ErrNotFound) {
			app.notFound(c)
			return
		}
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, apartment)
}

func (app *App) CreateApartmentHandler(c *gin.Context) {
	var req dto.CreateApartmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		app.clientError(c, http.StatusBadRequest, err.Error())
		return
	}

	response, err := app.ApartmentService.CreateApartment(c.Request.Context(), req)
	if err != nil {
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (app *App) UpdateApartmentHandler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		app.clientError(c, http.StatusBadRequest, "invalid apartment id format")
		return
	}

	var req dto.UpdateApartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		app.clientError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := app.ApartmentService.UpdateApartment(c.Request.Context(), id, req); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			app.notFound(c)
			return
		}
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "apartment updated successfully"})
}

func (app *App) DeleteApartmentHandler(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		app.clientError(c, http.StatusBadRequest, "invalid apartment id format")
		return
	}

	if err := app.ApartmentService.DeleteApartment(c.Request.Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			app.notFound(c)
			return
		}
		app.serverError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "apartment deleted successfully"})
}
