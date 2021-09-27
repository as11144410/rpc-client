package controllers

import (
	"github.com/gin-gonic/gin"
	"rpc-client/service"
	helloApi "rpc-service/api/protobuf-spec/helloworldpb"
)

func Hello(c *gin.Context) {
	req := helloApi.HelloRequest{
		Name: "haha",
	}
	res, _ := service.HelloService.SayHello(c, &req)
	result := map[string]interface{}{}

	result["data"] = res
	c.JSON(200, result)
}
