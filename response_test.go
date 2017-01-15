package webresponse_test

import (
	"testing"

	"github.com/stuwilli/go-test-utils"
	wr "github.com/stuwilli/go-web-response"
)

var testResponse wr.Response

func init() {

	testResponse = wr.Response{Message: "Test",
		StatusCode: 200,
		Timestamp:  1,
	}
}

func TestWriteJSON(t *testing.T) {

	data := make(map[string]string)
	data["test1"] = "foo"
	data["test2"] = "bar"

	testResponse.Data = data
	w := fakeresponse.NewFakeResponse(t)
	testResponse.WriteJSON(w)

	expected := `{"message":"Test","timestamp":1,"status":200,"data":{"test1":"foo","test2":"bar"},"path":""}`
	w.Assert(200, expected)
}
