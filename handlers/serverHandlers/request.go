package handlers

import "strings"

type Request struct {
	Method string
	Path   string
	Body   string
}

//Constructing Request
func NewRequest(request string) (req *Request, err bool) {

	lines := strings.Split(request, "\n")

	if len(lines) < 2 {
		return nil, true
	}
	firstLineParts := strings.Split(lines[0], " ")
	method := firstLineParts[0]
	path := firstLineParts[1]

	var body string

	for i := 0; i < len(request); i++ {
		if request[i] == '{' {
			for j := i; j < len(request); j++ {
				body += string(request[j])
			}
			break
		}
	}

	return &Request{
		Method: method,
		Path:   path,
		Body:   body,
	}, false
}
