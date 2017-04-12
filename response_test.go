package webresponse

import (
	"errors"
	"testing"

	"github.com/stuwilli/go-test-utils"
)

func TestResponseBuilder_WithDataStruct(t *testing.T) {

	type Test struct {
		Prop string
	}

	x := Test{Prop: "Hi"}

	b := NewBuilder()

	r := b.Status(200).Data(x).Build()

	r.Timestamp = 1

	if r.Data.(Test).Prop != "Hi" {
		t.Error("Expected hi")
	}

	expected := `{"statusCode":200,"message":"OK","timestamp":1,"data":{"Prop":"Hi"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(200, expected)

}

func TestResponseBuilder_WithRealError(t *testing.T) {

	b := NewBuilder()
	e1 := errors.New("An error")
	e2 := errors.New("Another error")
	r := b.Status(500).Error(e1).Error(e2).Build()
	r.Timestamp = 1
	//fmt.Println(string(r.marshalJSON()))

	expected := `{"statusCode":500,"message":"Internal Server Error","timestamp":1,"errors":{"error_1":"An error","error_2":"Another error"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(500, expected)

}

func TestResponseBuilder_WithNamedError(t *testing.T) {

	b := NewBuilder()
	r := b.Status(500).NamedError("error1", "An error").NamedError("error2", "Another error").Build()
	r.Timestamp = 1

	expected := `{"statusCode":500,"message":"Internal Server Error","timestamp":1,"errors":{"error1":"An error","error2":"Another error"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(500, expected)
}

func TestResponseBuilder_WithErrorsMap(t *testing.T) {
	err := map[string]string{
		"error1": "An error",
		"error2": "Another error",
	}

	b := NewBuilder()
	r := b.Status(500).Errors(err).Build()
	r.Timestamp = 1

	expected := `{"statusCode":500,"message":"Internal Server Error","timestamp":1,"errors":{"error1":"An error","error2":"Another error"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(500, expected)
}

func TestResponseBuilder_ErrorsWithError(t *testing.T) {

	b := NewBuilder()
	r := b.Status(500).Errors(errors.New("An error")).Build()
	r.Timestamp = 1

	expected := `{"statusCode":500,"message":"Internal Server Error","timestamp":1,"errors":{"error_1":"An error"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(500, expected)
}

func TestResponseBuilder_NewResponseData(t *testing.T) {

	type Test struct {
		Prop string
	}

	data := Test{Prop: "Hi"}

	r := NewResponse(200, data, nil)
	r.Timestamp = 1

	expected := `{"statusCode":200,"message":"OK","timestamp":1,"data":{"Prop":"Hi"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(200, expected)
}

func TestResponseBuilder_NewResponseError(t *testing.T) {

	r := NewResponse(500, nil, errors.New("An error"))
	r.Timestamp = 1

	expected := `{"statusCode":500,"message":"Internal Server Error","timestamp":1,"errors":{"error_1":"An error"}}`

	w := fakeresponse.NewFakeResponse(t)

	r.WriteJSON(w)
	w.Assert(500, expected)
}
