package gclient

import (
	"fmt"
	"log"
	"strings"
	
	"goshop/front-api/pkg/grpc/etcd3"
	"goshop/front-api/pkg/utils"
	
	"github.com/shinmigo/pb/memberpb"
	
	"github.com/shinmigo/pb/productpb"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	ProductClient         productpb.ProductServiceClient
	ProductCategoryClient productpb.CategoryServiceClient
	CartClient            memberpb.CartServiceClient
)

func DialGrpcService() {
	pms()
	crm()
}

func crm() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["crm"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["crm"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["crm"], strings.Join(utils.C.Etcd.Host, ","))
	CartClient = memberpb.NewCartServiceClient(conn)
}

func pms() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["pms"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["pms"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["pms"], strings.Join(utils.C.Etcd.Host, ","))
	ProductClient = productpb.NewProductServiceClient(conn)
	ProductCategoryClient = productpb.NewCategoryServiceClient(conn)
}
