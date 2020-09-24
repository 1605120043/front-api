package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/basepb"
	"github.com/shinmigo/pb/memberpb"
	"goshop/front-api/model/address"
	"goshop/front-api/pkg/grpc/gclient"
)

type Address struct {
	*gin.Context
}

func NewAddress(c *gin.Context) *Address {
	return &Address{Context: c}
}

func (m *Address) Index() ([]*address.AddressList, error) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	req := &memberpb.ListAddressReq{
		MemberId: memberId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	addressList, err := gclient.AddressClient.GetAddressListByMemberId(ctx, req)
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("获取列表失败, err:%v", err)
	}
	
	list := make([]*address.AddressList, 0, addressList.Total)
	if addressList.Total > 0 {
		for k := range addressList.Addresses {
			isDefault := false
			if addressList.Addresses[k].IsDefault == memberpb.AddressIsDefault_Used {
				isDefault = true
			}
			list = append(list, &address.AddressList{
				Id:        addressList.Addresses[k].AddressId,
				Name:      addressList.Addresses[k].Name,
				Mobile:    addressList.Addresses[k].Mobile,
				Address:   addressList.Addresses[k].Address,
				IsDefault: isDefault,
			})
		}
	}
	return list, err
}

func (m *Address) Detail() (*address.AddressDetail, error) {
	addressId, _ := strconv.ParseUint(m.DefaultQuery("address_id", "0"), 10, 64)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	addressDetail, err := gclient.AddressClient.GetAddressDetail(ctx, &basepb.GetOneReq{
		Id: addressId,
	})
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("获取详情失败, err:%v", err)
	}
	
	if addressDetail.MemberId != memberId {
		return nil, fmt.Errorf("获取详情失败, err:%v", err)
	}
	
	isDefault := false
	if addressDetail.IsDefault == memberpb.AddressIsDefault_Used {
		isDefault = true
	}
	
	detail := &address.AddressDetail{
		Name:      addressDetail.Name,
		Mobile:    addressDetail.Mobile,
		IsDefault: isDefault,
	}
	return detail, err
	
}

// 添加收货地址
func (m *Address) Add() error {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	longitude := m.PostForm("longitude")
	latitude := m.PostForm("latitude")
	
	isDef := 0
	if isDefault == "true" {
		isDef = 1
	}
	req := &memberpb.Address{
		MemberId:   memberId,
		Name:       name,
		Mobile:     mobile,
		Address:    addressDetail,
		RoomNumber: roomNumber,
		IsDefault:  memberpb.AddressIsDefault(isDef),
		Longitude:  longitude,
		Latitude:   latitude,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.AddressClient.AddAddress(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("添加收货地址失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}
	
	return nil
}

// 编辑收货地址
func (m *Address) Edit() error {
	addressId, _ := strconv.ParseUint(m.PostForm("id"), 10, 64)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	longitude := m.PostForm("longitude")
	latitude := m.PostForm("latitude")
	
	isDef := 0
	if isDefault == "true" {
		isDef = 1
	}
	req := &memberpb.Address{
		AddressId:  addressId,
		MemberId: memberId,
		Name:       name,
		Mobile:     mobile,
		Address:    addressDetail,
		RoomNumber: roomNumber,
		IsDefault:  memberpb.AddressIsDefault(isDef),
		Longitude:  longitude,
		Latitude:   latitude,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.AddressClient.EditAddress(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("更新收货地址失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("更新失败")
	}
	
	return nil
}
