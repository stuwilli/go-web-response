package webresponse

import (
	"encoding/json"
	"net/http"
)

//Response ...
type Response struct {
	Message    string      `json:"message"`
	Timestamp  int64       `json:"timestamp"`
	StatusCode int         `json:"status"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"errors,omitempty"`
	Path       string      `json:"path"`
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
