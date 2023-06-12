// Code generated by goa v3.11.3, DO NOT EDIT.
//
// calc gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/shibayu36/go-playground/diary/design

package client

import (
	"context"

	calc "github.com/shibayu36/go-playground/diary/gen/calc"
	calcpb "github.com/shibayu36/go-playground/diary/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildAddFunc builds the remote method to invoke for "calc" service "add"
// endpoint.
func BuildAddFunc(grpccli calcpb.CalcClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Add(ctx, reqpb.(*calcpb.AddRequest), opts...)
		}
		return grpccli.Add(ctx, &calcpb.AddRequest{}, opts...)
	}
}

// EncodeAddRequest encodes requests sent to calc add endpoint.
func EncodeAddRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*calc.AddPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "*calc.AddPayload", v)
	}
	return NewProtoAddRequest(payload), nil
}

// DecodeAddResponse decodes responses from the calc add endpoint.
func DecodeAddResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*calcpb.AddResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "*calcpb.AddResponse", v)
	}
	res := NewAddResult(message)
	return res, nil
}