package webresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Status ...
type Status struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//Response  ....
type Response struct {
	Status
	Timestamp int64             `json:"timestamp"`
	Data      interface{}       `json:"data,omitempty"`
	Errors    map[string]string `json:"errors,omitempty"`
	Pageable  *Pageable         `json:"pageable,omitempty"`
}

//Pageable ....
type Pageable struct {
	Current int `json:"current,omitempty"`
	Total   int `json:"total,omitempty"`
	PerPage int `json:"perPage,omitempty"`
}

//NewResponse ....
func NewResponse(statusCode int, data interface{}, err error) Response {

	status := Status{StatusCode: statusCode, Message: http.StatusText(statusCode)}
	resp := Response{Status: status, Timestamp: time.Now().Unix()}

	if data != nil {
		resp.Data = data
	}

	if err != nil {
		resp.AddError(err)
	}

	return resp
}

//AddError ....
func (r *Response) AddError(err error) {

	if r.Errors == nil {
		r.Errors = make(map[string]string)
	}

	count := len(r.Errors) + 1
	r.Errors[fmt.Sprintf("error_%v", count)] = err.Error()
}

//AddNamedError ...
func (r *Response) AddNamedError(name string, err error) {

	if r.Errors == nil {
		r.Errors = make(map[string]string)
	}

	r.Errors[name] = err.Error()
}

//MarshalJSON ...
func (r *Response) marshalJSON() []byte {

	out, _ := json.Marshal(r)

	return out
}

//WriteJSON ...
func (r *Response) WriteJSON(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	w.Write(r.marshalJSON())
}
