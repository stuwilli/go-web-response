package webresponse

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stuwilli/go-test-utils"
)

func TestNewResponse_WithData(t *testing.T) {

	data := map[string]string{
		"Test1": "This is test one",
		"Test2": "This is test two",
	}

	r := NewResponse(200, data, nil)

	if len(r.Data.(map[string]string)) != 2 {
		t.Error("Expected 2 data attributes")
	}

	if r.Message != "OK" {
		t.Error("Expected message to be 'OK', got", r.Message)
	}
}

func TestNewResponse_WithError(t *testing.T) {

	r := NewResponse(http.StatusBadRequest, nil, errors.New("Test"))

	if len(r.Errors) != 1 {
		t.Error("Expected there to be 1 error")
	}
}

//TestAddError ...
func TestAddError(t *testing.T) {

	r := NewResponse(200, nil, nil)

	r.AddError(errors.New("Test"))

	if r.Errors["error_1"] != "Test" {
		t.Error("Expected error_1 to be set and to equal Test")
	}

	r.AddError(errors.New("Test2"))

	if len(r.Errors) != 2 {
		t.Error("Expected error count to be 2, got", len(r.Errors))
	}

	if r.Errors["error_2"] != "Test2" {
		t.Error("Expected error_2 to be set and to equal Test2")
	}
}

func TestAddNamedError(t *testing.T) {

	r := NewResponse(200, nil, nil)

	r.AddNamedError("test", errors.New("Test"))

	if len(r.Errors) != 1 {
		t.Error("Expected one error")
	}

	if r.Errors["test"] != "Test" {
		t.Error("Expected error with key 'test' and value 'Test'")
	}
}

func TestWriteJSON(t *testing.T) {

	data := map[string]string{
		"Test1": "This is test one",
		"Test2": "This is test two",
	}

	r := NewResponse(200, data, errors.New("Test"))
	r.Timestamp = 1
	w := fakeresponse.NewFakeResponse(t)
	r.WriteJSON(w)
	equal := `{"statusCode":200,"message":"OK","timestamp":1,"data":{"Test1":"This is test one","Test2":"This is test two"},"errors":{"error_1":"Test"}}`
	w.Assert(200, equal)
}
