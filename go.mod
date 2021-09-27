module rpc-client

go 1.15

require (
	github.com/gin-gonic/gin v1.7.4
	google.golang.org/grpc v1.41.0
	rpc-service v0.0.0
)

replace rpc-service => ../rpc-service
