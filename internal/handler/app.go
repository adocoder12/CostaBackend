package handler

import (
	service "github.com/adocoder12/Costabackend/internal/services"
	"log"
	"net/http"
)

type App struct {
	InfoLog          *log.Logger
	ErrorLog         *log.Logger
	ApartmentService *service.ApartmentService
}

func NewServer(infoLog *log.Logger, errorLog *log.Logger) *App {
	return &App{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
func (server *App) serverError(w http.ResponseWriter, err error) {
	server.ErrorLog.Printf("%+v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (server *App) clientError(w http.ResponseWriter, status int) {
	// Standardized way to send 400, 401, 403, etc.
	http.Error(w, http.StatusText(status), status)
}
func (server *App) notFound(w http.ResponseWriter) {
	server.clientError(w, http.StatusNotFound)
}
