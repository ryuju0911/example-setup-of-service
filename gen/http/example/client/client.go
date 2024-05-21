// Code generated by goa v3.16.2, DO NOT EDIT.
//
// example client HTTP transport
//
// Command:
// $ goa gen example/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the example service endpoint HTTP clients.
type Client struct {
	// CreateSomeData Doer is the HTTP client used to make requests to the
	// CreateSomeData endpoint.
	CreateSomeDataDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the example service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CreateSomeDataDoer:  doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// CreateSomeData returns an endpoint that makes HTTP requests to the example
// service CreateSomeData server.
func (c *Client) CreateSomeData() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateSomeDataRequest(c.encoder)
		decodeResponse = DecodeCreateSomeDataResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCreateSomeDataRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateSomeDataDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("example", "CreateSomeData", err)
		}
		return decodeResponse(resp)
	}
}
