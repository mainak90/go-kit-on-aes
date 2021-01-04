package helpers

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EncryptService interface {
	Encrypt(context.Context, string, string) (string, error)
	Decrypt(context.Context, string, string) (string, error)
	GenKey(context.Context, int)(string, error)
}

func MakeEncryptEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(EncryptRequest)
		message, err := svc.Encrypt(ctx, req.Key, req.Text)
		if err != nil {
			return EncryptResponse{"", err.Error()}, nil
		}
		return EncryptResponse{message, ""}, nil
	}
}

func MakeDecryptEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DecryptRequest)
		text, err := svc.Decrypt(ctx, req.Message, req.Key)
		if err != nil {
			return DecryptResponse{"", err.Error()}, nil
		}
		return DecryptResponse{text, ""}, nil
	}
}

func MakeGenKeyEndpoint(svc EncryptService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var keybit = request.(GenKeyRequest)
		key, err := svc.GenKey(ctx, keybit.Bit)
		if err != nil {
			return GenKeyResponse{"", err.Error()}, nil
		}
		return GenKeyResponse{key, ""}, nil
	}
}