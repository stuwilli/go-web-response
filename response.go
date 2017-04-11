package webresponse

import (
	"encoding/json"
	"net/http"
)

//Status ...
type Status struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//Response  ....
// type Response interface {
// 	WriteJSON(w http.ResponseWriter)
// 	marshalJSON() []byte
// }

//ResponseBuilder ...
type ResponseBuilder interface {
	SetStatus(int) ResponseBuilder
	SetData(interface{}) ResponseBuilder
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
	Status int
	Data   interface{}
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

func (r *responseBuilder) SetData(d interface{}) ResponseBuilder {
	r.Data = d
	return r
}

func (r *responseBuilder) SetStatus(s int) ResponseBuilder {
	r.Status = s
	return r
}

func (r *responseBuilder) Build() Response {

	status := Status{StatusCode: r.Status, Message: http.StatusText(r.Status)}
	return Response{Status: status, Data: r.Data}
}
