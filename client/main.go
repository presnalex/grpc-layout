package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/presnalex/go-micro/v3/service"
	"github.com/presnalex/go-micro/v3/wrapper/requestid"
	pb "github.com/presnalex/grpc-layout/grpc-layout-proto/go/proto"
	mclient "go.unistack.org/micro-client-grpc/v3"
	"go.unistack.org/micro/v3/client"
	"go.unistack.org/micro/v3/metadata"
)

func main() {
	cli1opts, _ := service.ClientOptions(&service.ClientConfig{})
	cli1opts = append(cli1opts, client.Retries(0), client.ContentType("application/grpc+proto"))

	c := mclient.NewClient(cli1opts...)

	c = client.NewClientCallOptions(c, client.WithAddress("localhost:9090"))

	client := pb.NewPingPongService("pingpong.grpc.client", c)

	//Пример создания контекста с X-Request-Id
	uid, err := uuid.NewRandom()
	if err != nil {
		uid = uuid.Nil
	}
	md := metadata.New(2)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	ctx = requestid.SetOutgoingRequestId(ctx, uid.String())

	rsp, err := client.Call(ctx, &pb.RequestMsg{Msg: "ping"})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v\n\n", rsp)

}
