// Code generated by zanzibar
// @generated

package bar_test

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/test/lib/test_gateway"
)

// TODO(zw): benchBytes should be generated according to request type.
var benchBytesForNormal = []byte("{}")

func TestNormalOKResponse(t *testing.T) {
	var counter int = 0

	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		KnownBackends: []string{"bar"},
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	fakeNormal := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		// TODO(zw): generate client response.
		if _, err := w.Write([]byte("{}")); err != nil {
			t.Fatal("can't write fake response")
		}
		counter++
	}
	gateway.Backends()["bar"].HandleFunc(
		"POST", "/bar-path", fakeNormal,
	)

	res, err := gateway.MakeRequest(
		"POST", "/bar/bar-path", bytes.NewReader(benchBytesForNormal),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, "200 OK", res.Status)
	assert.Equal(t, 1, counter)
}
