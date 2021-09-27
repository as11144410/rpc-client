package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"rpc-client/controllers"
	"rpc-client/service"
	"time"
)

const (
	address     = "localhost:8080"
	defaultName = "world"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	rand.Seed(time.Now().UnixNano())

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/hello", controllers.Hello)
	defer service.GrpcCnn.Close()
	fmt.Printf("%s", r.Run(":8081"))
}
