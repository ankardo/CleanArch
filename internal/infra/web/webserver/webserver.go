package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Route struct {
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string][]Route
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string][]Route),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.Handlers[path] = append(s.Handlers[path], Route{
		Method:  method,
		Handler: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)

	for path, routes := range s.Handlers {
		for _, route := range routes {
			switch route.Method {
			case http.MethodGet:
				s.Router.Get(path, route.Handler)
			case http.MethodPost:
				s.Router.Post(path, route.Handler)
			}
		}
	}

	http.ListenAndServe(s.WebServerPort, s.Router)
}
