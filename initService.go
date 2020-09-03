package main

import (
	"goshop/front-api/pkg/grpc/gclient"
)

func initService() {
	go gclient.DialGrpcService()
}
