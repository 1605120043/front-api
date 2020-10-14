package main

import (
	"goshop/front-api/command/payment"
	"goshop/front-api/pkg/grpc/gclient"
)

func initService() {
	go gclient.DialGrpcService()
	go payment.GetPaymentConfig()
}
