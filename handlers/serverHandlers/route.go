package handlers

type Route struct {
	Path   string
	Method string
}

// Constructing Route
func NewRoute(path string, method string) *Route {
	return &Route{
		Path:   path,
		Method: method,
	}
}

// Signature and func reference for handling request with router
type RequestHandler func(param string) string
