package config

import "rpc-client/utils"

var HostDomain = utils.GetEvnWithDefaultVal("HOST_DOMAIN", "https://onepagedev.ngrok.starservices.store")

var APIVersion = utils.GetEvnWithDefaultVal("API_VERSION", "2021-01")

var GRPC_SERVER_ADDR = utils.GetEvnWithDefaultVal("GRPC_SERVER_ADDR", "127.0.0.1:8080")
