package handlers

import "fmt"

type Response struct {
	Status  int
	Message string
}

//Constructing Response
func NewResponse() *Response {
	return &Response{}
}

//Formatting response string into format supported by http
func (r *Response) makeResponse(status string, data string) string {
	return fmt.Sprintf("%v\r\nContent-Lenth: %v\r\nAccess-Control-Allow-Origin: *\r\n\r\n%v", status, len(data), data)
}

//200 OK http response
func (r *Response) Ok(data string) string {
	status := "HTTP/1.1 200 OK"

	return r.makeResponse(status, data)
}

//201 Created http response
func (r *Response) Created(data string) string {
	status := "HTTP/1.1 201 Created"

	return r.makeResponse(status, data)
}

//400 Bad request http response
func (r *Response) BadRequest(data string) string {
	status := "HTTP/1.1 400 Bad Request"

	return fmt.Sprintf("%v\r\n\r\n", status)
}

//404 not found http response
func (r *Response) NotFound(data string) string {
	status := "HTTP/1.1 404 Not Found"

	return r.makeResponse(status, data)
}
