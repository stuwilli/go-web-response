package webresponse

import (
	"fmt"
	"testing"
)

func TestNewResponse_WithData(t *testing.T) {

	type Test struct {
		Prop string
	}

	x := Test{Prop: "Hi"}

	b := NewBuilder()

	r := b.SetStatus(200).SetData(x).Build()

	fmt.Println(string(r.marshalJSON()))

	if r.Data.(Test).Prop != "Hi" {
		t.Error("Expected hi")
	}
}
