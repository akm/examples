// Code generated by goa v3.0.9, DO NOT EDIT.
//
// text client
//
// Command:
// $ goa gen goa.design/examples/encodings/design -o
// $(GOPATH)/src/goa.design/examples/encodings

package text

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "text" service client.
type Client struct {
	ConcatstringsEndpoint     goa.Endpoint
	ConcatbytesEndpoint       goa.Endpoint
	ConcatstringfieldEndpoint goa.Endpoint
	ConcatbytesfieldEndpoint  goa.Endpoint
}

// NewClient initializes a "text" service client given the endpoints.
func NewClient(concatstrings, concatbytes, concatstringfield, concatbytesfield goa.Endpoint) *Client {
	return &Client{
		ConcatstringsEndpoint:     concatstrings,
		ConcatbytesEndpoint:       concatbytes,
		ConcatstringfieldEndpoint: concatstringfield,
		ConcatbytesfieldEndpoint:  concatbytesfield,
	}
}

// Concatstrings calls the "concatstrings" endpoint of the "text" service.
func (c *Client) Concatstrings(ctx context.Context, p *ConcatstringsPayload) (res string, err error) {
	var ires interface{}
	ires, err = c.ConcatstringsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}

// Concatbytes calls the "concatbytes" endpoint of the "text" service.
func (c *Client) Concatbytes(ctx context.Context, p *ConcatbytesPayload) (res []byte, err error) {
	var ires interface{}
	ires, err = c.ConcatbytesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]byte), nil
}

// Concatstringfield calls the "concatstringfield" endpoint of the "text"
// service.
func (c *Client) Concatstringfield(ctx context.Context, p *ConcatstringfieldPayload) (res *MyConcatenation, err error) {
	var ires interface{}
	ires, err = c.ConcatstringfieldEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MyConcatenation), nil
}

// Concatbytesfield calls the "concatbytesfield" endpoint of the "text" service.
func (c *Client) Concatbytesfield(ctx context.Context, p *ConcatbytesfieldPayload) (res *MyConcatenation, err error) {
	var ires interface{}
	ires, err = c.ConcatbytesfieldEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*MyConcatenation), nil
}
