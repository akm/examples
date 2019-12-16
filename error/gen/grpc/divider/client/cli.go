// Code generated by goa v3.0.9, DO NOT EDIT.
//
// divider gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"encoding/json"
	"fmt"

	divider "goa.design/examples/error/gen/divider"
	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
)

// BuildIntegerDividePayload builds the payload for the divider integer_divide
// endpoint from CLI flags.
func BuildIntegerDividePayload(dividerIntegerDivideMessage string) (*divider.IntOperands, error) {
	var err error
	var message dividerpb.IntegerDivideRequest
	{
		if dividerIntegerDivideMessage != "" {
			err = json.Unmarshal([]byte(dividerIntegerDivideMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"a\": 1338266005399228665,\n      \"b\": 7245195139064803075\n   }'")
			}
		}
	}
	v := &divider.IntOperands{
		A: int(message.A),
		B: int(message.B),
	}
	return v, nil
}

// BuildDividePayload builds the payload for the divider divide endpoint from
// CLI flags.
func BuildDividePayload(dividerDivideMessage string) (*divider.FloatOperands, error) {
	var err error
	var message dividerpb.DivideRequest
	{
		if dividerDivideMessage != "" {
			err = json.Unmarshal([]byte(dividerDivideMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, example of valid JSON:\n%s", "'{\n      \"a\": 0.7822693555171186,\n      \"b\": 0.5749246657891343\n   }'")
			}
		}
	}
	v := &divider.FloatOperands{
		A: message.A,
		B: message.B,
	}
	return v, nil
}
