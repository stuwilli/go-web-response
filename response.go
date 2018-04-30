//Package webresponse provides a builder for creating JSON and XML web responses
package webresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Status holds the status code and message
type Status struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

//ResponseBuilder interface provides the builder methods
type ResponseBuilder interface {
	Status(int) ResponseBuilder
	Data(interface{}) ResponseBuilder
	Errors(interface{}) ResponseBuilder
	Error(error) ResponseBuilder
	NamedError(string, string) ResponseBuilder
	Build() Response
}

//Response holds the reponse data
type Response struct {
	Status
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
}

//responseBuilder impliments ResponseBuilder
type responseBuilder struct {
	status int
	data   interface{}
	errors interface{}
}

//NewBuilder creates a new ResponseBuilder, the status code defaults to 200
func NewBuilder() ResponseBuilder {
	return &responseBuilder{status: 200}
}

//MarshalJSON marshalls the Response to JSON and returns a byte array
func (r *Response) marshalJSON() []byte {

	out, _ := json.Marshal(r)

	return out
}

//WriteJSON writes the JSON response to a response writer w
func (r *Response) WriteJSON(w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	w.Write(r.marshalJSON())
}

//Data adds a data item to the reponse can be anything and returns the ResponseBuilder
func (r *responseBuilder) Data(d interface{}) ResponseBuilder {
	r.data = d
	return r
}

//Status sets the status code for the response and return the ResponseBuilder
func (r *responseBuilder) Status(s int) ResponseBuilder {
	r.status = s
	return r
}

//Errors adds a custom error type or a single error and returns ResponseBuilder
func (r *responseBuilder) Errors(e interface{}) ResponseBuilder {

	switch e.(type) {

	case error:
		r.Error(e.(error))

	default:
		r.errors = e
	}

	return r
}

//checkAndInitErrors checks to see if errors has been initialised
//if not it initialises a map[string]string and returns the index +1
//if already initialised as a map[string]string it return the current index + 1
func (r *responseBuilder) checkAndInitErrors() int {

	_, ok := r.errors.(map[string]string)

	if r.errors == nil || !ok {
		r.errors = make(map[string]string)
	}

	return len(r.errors.(map[string]string)) + 1
}

//Error takes an error and adds the message to a map[string]string
//returns ResponseBuilder
func (r *responseBuilder) Error(e error) ResponseBuilder {

	count := r.checkAndInitErrors()

	r.NamedError(fmt.Sprintf("error_%v", count), e.Error())

	return r
}

//NamedError add a named error to the error map returns the ResponseBuilder
func (r *responseBuilder) NamedError(k string, v string) ResponseBuilder {

	r.checkAndInitErrors()

	r.errors.(map[string]string)[k] = v

	return r
}

//Build creates and return the Response
func (r *responseBuilder) Build() Response {

	status := Status{StatusCode: r.status, Message: http.StatusText(r.status)}
	return Response{Status: status, Data: r.data, Errors: r.errors,
		Timestamp: makeTimestamp()}
}

//NewResponse legacy response building method
func NewResponse(status int, data interface{}, err error) Response {

	b := NewBuilder().Status(status)

	if data != nil {
		b.Data(data)
	}

	if err != nil {
		b.Error(err)
	}

	return b.Build()
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
