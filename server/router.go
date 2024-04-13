package server

import (
	handlers "github/perasd9/MTWebServer/handlers/serverHandlers"
	"regexp"
)

type Router struct {
	mappings map[handlers.Route]handlers.RequestHandler
}

// Construcing router
func NewRouter() *Router {

	return &Router{
		mappings: make(map[handlers.Route]handlers.RequestHandler),
	}
}

// Finding out which route is supported by request
func (r *Router) Handle(request *handlers.Request) (string, bool) {
	// path := request.Path
	route := *handlers.NewRoute(request.Path, request.Method)

	// handler, ok := r.mappings[route]
	// if !ok {
	// 	fmt.Println("No handler found for path")
	// 	return "", false
	// }

	for key, value := range r.mappings {
		match, err := regexp.MatchString("\\"+key.Path+"\\b", route.Path)
		if err != nil {
			continue
		}
		if match && key.Method == route.Method {

			if request.Body != "" {
				return value(request.Body), true
			} else {
				return value(""), true
			}
		}
	}

	//Calling actual handler support for route
	return handlers.NewResponse().NotFound(""), false
}

// Adding route to all defined routes of our API
func (r *Router) AddRoute(route handlers.Route, handler handlers.RequestHandler) {
	r.mappings[route] = handler
}
