// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"

	clientTypeBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/github.com/uber/zanzibar/clients/bar/bar"
)

// HandleNormalRequest handles "/bar/bar-path".
func HandleNormalRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {
	// Handle request headers.
	h := http.Header{}

	// Handle request body.
	var body NormalHTTPRequest
	if ok := req.ReadAndUnmarshalBody(&body); !ok {
		return
	}
	clientRequest := convertToNormalClientRequest(&body)
	clientResp, err := clients.Bar.Normal(ctx, clientRequest, h)
	if err != nil {
		req.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		res.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	defer func() {
		if cerr := clientResp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Handle client respnse.
	if !res.IsOKResponse(clientResp.StatusCode, []int{200}) {
		req.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	b, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		res.SendError(500, errors.Wrap(err, "could not read client response body:"))
		return
	}
	var clientRespBody bar.BarResponse
	if err := clientRespBody.UnmarshalJSON(b); err != nil {
		res.SendError(500, errors.Wrap(err, "could not unmarshal client response body:"))
		return
	}
	response := convertNormalClientResponse(&clientRespBody)
	res.WriteJSON(clientResp.StatusCode, response)
	res.Flush()
}

func convertToNormalClientRequest(body *NormalHTTPRequest) *barClient.NormalHTTPRequest {
	clientRequest := &barClient.NormalHTTPRequest{}

	clientRequest.Request = clientTypeBar.BarRequest(body.Request)

	return clientRequest
}
func convertNormalClientResponse(body *bar.BarResponse) *bar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := &bar.BarResponse{}
	return downstreamResponse
}
