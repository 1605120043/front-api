package gclient

import (
	"fmt"
	"log"
	"strings"
	
	"github.com/shinmigo/pb/orderpb"
	
	"goshop/front-api/pkg/grpc/etcd3"
	"goshop/front-api/pkg/utils"
	
	"github.com/shinmigo/pb/shoppb"
	
	"github.com/shinmigo/pb/memberpb"
	
	"github.com/shinmigo/pb/productpb"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	ProductClient         productpb.ProductServiceClient
	ProductCategoryClient productpb.CategoryServiceClient
	CartClient            memberpb.CartServiceClient
	AddressClient         memberpb.AddressServiceClient
	MemberClient          memberpb.MemberServiceClient
	AreaClient            shoppb.AreaServiceClient
	OrderClient           orderpb.OrderServiceClient
	TagClient             productpb.TagServiceClient
)

func DialGrpcService() {
	pms()
	crm()
	shop()
	oms()
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
	AddressClient = memberpb.NewAddressServiceClient(conn)
	MemberClient = memberpb.NewMemberServiceClient(conn)
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
	TagClient = productpb.NewTagServiceClient(conn)
}

func shop() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["shop"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["shop"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["shop"], strings.Join(utils.C.Etcd.Host, ","))
	AreaClient = shoppb.NewAreaServiceClient(conn)
}

func oms() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["oms"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["oms"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["oms"], strings.Join(utils.C.Etcd.Host, ","))
	OrderClient = orderpb.NewOrderServiceClient(conn)
}
