package handler

import (
	"context"

	pb "github.com/presnalex/grpc-layout/grpc-layout-proto/go/proto"
)

const (
	noErrorStatus      = "0"
	dataNotFoundStatus = "-1"
)

type Handler struct {
}

func (e *Handler) Call(ctx context.Context, req *pb.RequestMsg, rsp *pb.ResponseMsg) error {
	rsp.Msg = "pong"
	return nil
}
