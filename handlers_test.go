package webresponse_test

import (
	"errors"
	"testing"

	wr "github.com/stuwilli/go-web-response"
	validate "github.com/stuwilli/govalidate"
)

func TestNewOkResponse(t *testing.T) {

	data := make(map[string]string)
	data["test1"] = "foo"
	data["test2"] = "bar"

	path := "/service"

	res := wr.NewOkResponse(data, path)

	if res.Message != "Ok" {
		t.Error("Expected Message to equal Ok, got,", res.Message)
	}

	if res.StatusCode != 200 {
		t.Error("Expected StatusCode to be 200, got", res.StatusCode)
	}

	if res.Data.(map[string]string)["test1"] != "foo" {
		t.Error("Expected Data.test1 to equal foo, got", res.Data.(map[string]string)["test1"])
	}
}

func TestNewCreatedResponse(t *testing.T) {

	data := make(map[string]string)
	data["test1"] = "foo"
	data["test2"] = "bar"

	path := "/service"

	res := wr.NewCreatedResponse(data, path)

	if res.Message != "Created" {
		t.Error("Expected Message to equal Ok, got,", res.Message)
	}

	if res.StatusCode != 201 {
		t.Error("Expected StatusCode to be 200, got", res.StatusCode)
	}

	if res.Data.(map[string]string)["test1"] != "foo" {
		t.Error("Expected Data.test1 to equal foo, got", res.Data.(map[string]string)["test1"])
	}
}

func TestNewNotFoundResponse_Validation(t *testing.T) {

	errs := make(map[string]string)
	errs["err1"] = "error1"
	errs["err2"] = "error2"
	err := validate.ValidationError{Err: errs}

	path := "/service"

	res := wr.NewNotFoundResponse(err, path)

	if res.StatusCode != 404 {
		t.Error("Expected StatusCode to be 200, got", res.StatusCode)
	}

	if res.Error.(map[string]string)["error"]["err1"] != errs["err1"] {
		t.Error("Expected Error to equal Test Error, got", res.Error.(map[string]string)["error"])
	}
}

//TODO TestNewNotFoundResponse_MySQL

func TestNewNotFoundResponse_String(t *testing.T) {

	err := "Test error"

	path := "/service"

	res := wr.NewNotFoundResponse(err, path)

	if res.Message != "NotFound" {
		t.Error("Expected Message to equal Ok, got,", res.Message)
	}

	if res.StatusCode != 404 {
		t.Error("Expected StatusCode to be 200, got", res.StatusCode)
	}

	if res.Error.(map[string]string)["error"] != "Test error" {
		t.Error("Expected Error to equal Test Error, got", res.Error.(map[string]string)["error"])
	}
}

func TestNewNotFoundResponse_Error(t *testing.T) {

	err := errors.New("Test error")

	path := "/service"

	res := wr.NewNotFoundResponse(err, path)

	if res.Message != "NotFound" {
		t.Error("Expected Message to equal Ok, got,", res.Message)
	}

	if res.StatusCode != 404 {
		t.Error("Expected StatusCode to be 200, got", res.StatusCode)
	}

	if res.Error.(map[string]string)["error"] != "Test error" {
		t.Error("Expected Error to equal Test Error, got", res.Error.(map[string]string)["error"])
	}
}

func TestNewNotFoundResponse_Unknown(t *testing.T) {

	path := "/service"

	res := wr.NewNotFoundResponse(nil, path)

	if res.Message != "NotFound" {
		t.Error("Expected Message to equal Ok, got,", res.Message)
	}

	if res.StatusCode != 404 {
		t.Error("Expected StatusCode to be 200, got", res.StatusCode)
	}

	if res.Error.(map[string]string)["error"] != "Something went wrong" {
		t.Error("Expected Error to equal Test Error, got", res.Error.(map[string]string)["error"])
	}
}
