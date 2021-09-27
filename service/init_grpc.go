package service

import (
	"context"
	"rpc-client/config"

	helloApi "rpc-service/api/protobuf-spec/helloworldpb"

	"log"

	"google.golang.org/grpc"
)

var GrpcCnn *grpc.ClientConn

var HelloService helloApi.GreeterClient

func init() {
	var err error
	GrpcCnn, err = getClientConn(context.Background(), config.GRPC_SERVER_ADDR, []grpc.DialOption{})
	if err != nil {
		log.Fatalf("Dial failed:%v", err)
	}

	HelloService = helloApi.NewGreeterClient(GrpcCnn)

	log.Println("completed grpc service !!!")

}

func getClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
