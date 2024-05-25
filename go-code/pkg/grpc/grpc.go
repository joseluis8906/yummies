package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/metadata"
)

func Header(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("metadata not found")
	}

	val := md.Get(key)
	if len(val) == 0 {
		return "", errors.New("header not found")
	}

	return val[0], nil
}
