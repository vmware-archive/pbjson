package pbjson_test

import (
	"testing"

	"github.com/bitnami-labs/pbjson"
	"github.com/bitnami-labs/pbjson/pbjsontest"
)

func TestDefaults(t *testing.T) {
	ex := &pbjsontest.TestDefaults{}
	b, err := pbjson.Marshal(ex, pbjson.WithDefaults(true))
	if err != nil {
		t.Fatal(err)
	}
	if got, want := string(b), `{"anExample":""}`; got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}
