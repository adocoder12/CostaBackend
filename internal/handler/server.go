package handler

import (
	"log"
	"net/http"
)

type Server struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func NewServer(infoLog *log.Logger, errorLog *log.Logger) *Server {
	return &Server{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
}
func (server *Server) serverError(w http.ResponseWriter, err error) {
	server.ErrorLog.Printf("%+v", err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (server *Server) clientError(w http.ResponseWriter, status int) {
	// Standardized way to send 400, 401, 403, etc.
	http.Error(w, http.StatusText(status), status)
}
func (server *Server) notFound(w http.ResponseWriter) {
	server.clientError(w, http.StatusNotFound)
}
