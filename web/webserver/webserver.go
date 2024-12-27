package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	handlers      map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer( port string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		handlers:      make(map[string]http.HandlerFunc),
		WebServerPort: port,
	}
}

func (s *WebServer) AddHandler(name string, handler http.HandlerFunc) {
	s.handlers[name] = handler
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for name, handler := range s.handlers {
		s.Router.Post(name, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)

}
