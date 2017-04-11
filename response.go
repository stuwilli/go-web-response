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

//ResponseBuilder ...
type ResponseBuilder interface {
	Status(int) ResponseBuilder
	Data(interface{}) ResponseBuilder
	Errors(interface{}) ResponseBuilder
	Error(error) ResponseBuilder
	NamedError(string, string) ResponseBuilder
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
	errors interface{}
}

//NewBuilder ...
func NewBuilder() ResponseBuilder {
	return &responseBuilder{status: 200}
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

func (r *responseBuilder) Errors(e interface{}) ResponseBuilder {

	switch e.(type) {

	case error:
		r.Error(e.(error))

	default:
		r.errors = e
	}

	return r
}

func (r *responseBuilder) checkAndInitErrors() int {

	_, ok := r.errors.(map[string]string)

	if r.errors == nil || !ok {
		r.errors = make(map[string]string)
	}

	return len(r.errors.(map[string]string)) + 1
}

func (r *responseBuilder) Error(e error) ResponseBuilder {

	count := r.checkAndInitErrors()

	r.NamedError(fmt.Sprintf("error_%v", count), e.Error())

	return r
}

func (r *responseBuilder) NamedError(k string, v string) ResponseBuilder {

	r.checkAndInitErrors()

	r.errors.(map[string]string)[k] = v

	return r
}

func (r *responseBuilder) Build() Response {

	status := Status{StatusCode: r.status, Message: http.StatusText(r.status)}
	return Response{Status: status, Data: r.data, Errors: r.errors,
		Timestamp: time.Now().Unix()}
}
