package webresponse

import (
	"encoding/json"
	"net/http"
	"time"
)

//Status ...
type Status struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//ResponseBuilder ...
type ResponseBuilder interface {
	Status(int) ResponseBuilder
	Data(interface{}) ResponseBuilder
	Build() Response
}

//Response ...
type Response struct {
	Status
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
}

type responseBuilder struct {
	status int
	data   interface{}
	Errors interface{}
}

//NewBuilder ...
func NewBuilder() ResponseBuilder {
	return &responseBuilder{}
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

func (r *responseBuilder) Data(d interface{}) ResponseBuilder {
	r.data = d
	return r
}

func (r *responseBuilder) Status(s int) ResponseBuilder {
	r.status = s
	return r
}

func (r *responseBuilder) Build() Response {

	status := Status{StatusCode: r.status, Message: http.StatusText(r.status)}
	return Response{Status: status, Data: r.data, Timestamp: time.Now().Unix()}
}
