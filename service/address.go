package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/basepb"
	"github.com/shinmigo/pb/memberpb"
	"github.com/shinmigo/pb/shoppb"
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
	
	if err != nil {
		return nil, fmt.Errorf("获取列表失败, err:%v", err)
	}
	
	list := make([]*address.AddressList, 0, addressList.Total)
	areaCodes := make([]uint64, 0, addressList.Total*3)
	if addressList.Total > 0 {
		for k := range addressList.Addresses {
			areaCodes = append(areaCodes, addressList.Addresses[k].CodeProv)
			areaCodes = append(areaCodes, addressList.Addresses[k].CodeCity)
			areaCodes = append(areaCodes, addressList.Addresses[k].CodeCoun)
		}
		// 根据code查找区域
		areaNameList, err := gclient.AreaClient.GetAreaNameByCodes(ctx, &shoppb.AreaCodeReq{Codes: areaCodes})
		if err != nil {
			return nil, fmt.Errorf("获取区域失败, err:%v", err)
		}
		
		areaNameByCode := make(map[uint64]string, len(areaNameList.Codes))
		for _, item := range areaNameList.Codes {
			areaNameByCode[item.Code] = item.Name
		}
		
		for k := range addressList.Addresses {
			spew.Dump(addressList.Addresses)
			
			isDefault := false
			if addressList.Addresses[k].IsDefault == memberpb.AddressIsDefault_Used {
				isDefault = true
			}
			list = append(list, &address.AddressList{
				AddressId: addressList.Addresses[k].AddressId,
				Name:      addressList.Addresses[k].Name,
				Mobile:    addressList.Addresses[k].Mobile,
				Address: fmt.Sprint(
					areaNameByCode[addressList.Addresses[k].CodeProv],
					areaNameByCode[addressList.Addresses[k].CodeCity],
					areaNameByCode[addressList.Addresses[k].CodeCoun],
					addressList.Addresses[k].Address,
					addressList.Addresses[k].RoomNumber,
				),
				IsDefault: isDefault,
			})
		}
	}
	cancel()
	
	return list, err
}

func (m *Address) Detail() (*address.AddressDetail, error) {
	addressId, _ := strconv.ParseUint(m.DefaultQuery("address_id", "0"), 10, 64)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	addressDetail, err := gclient.AddressClient.GetAddressDetail(ctx, &basepb.GetOneReq{
		Id: addressId,
	})
	
	if err != nil {
		return nil, fmt.Errorf("获取详情失败, err:%v", err)
	}
	
	if addressDetail.MemberId != memberId {
		return nil, fmt.Errorf("获取详情失败, err:%v", err)
	}
	
	areaCodes := []uint64{addressDetail.CodeProv, addressDetail.CodeCity, addressDetail.CodeCoun}
	// 根据code查找区域
	areaNameList, err := gclient.AreaClient.GetAreaNameByCodes(ctx, &shoppb.AreaCodeReq{Codes: areaCodes})
	if err != nil {
		return nil, fmt.Errorf("获取区域失败, err:%v", err)
	}
	areaNameByCode := make(map[uint64]string, len(areaNameList.Codes))
	for _, item := range areaNameList.Codes {
		areaNameByCode[item.Code] = item.Name
	}
	
	isDefault := false
	if addressDetail.IsDefault == memberpb.AddressIsDefault_Used {
		isDefault = true
	}
	
	detail := &address.AddressDetail{
		AddressId: addressId,
		Name:      addressDetail.Name,
		Mobile:    addressDetail.Mobile,
		CodeProv:  addressDetail.CodeProv,
		CodeCity:  addressDetail.CodeCity,
		CodeCoun:  addressDetail.CodeCoun,
		
		CodeProvName: areaNameByCode[addressDetail.CodeProv],
		CodeCityName: areaNameByCode[addressDetail.CodeCity],
		CodeCounName: areaNameByCode[addressDetail.CodeCoun],
		
		Address:    addressDetail.Address,
		RoomNumber: addressDetail.RoomNumber,
		IsDefault:  isDefault,
	}
	cancel()
	return detail, err
	
}

// 添加收货地址
func (m *Address) Add() error {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	codeProv, _ := strconv.ParseUint(m.PostForm("code_prov"), 10, 64)
	codeCity, _ := strconv.ParseUint(m.PostForm("code_city"), 10, 64)
	codeCoun, _ := strconv.ParseUint(m.PostForm("code_coun"), 10, 64)
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	//longitude := m.PostForm("longitude")
	//latitude := m.PostForm("latitude")
	
	isDef := 0
	if isDefault == "true" {
		isDef = 1
	}
	req := &memberpb.Address{
		MemberId:   memberId,
		Name:       name,
		Mobile:     mobile,
		CodeProv:   codeProv,
		CodeCity:   codeCity,
		CodeCoun:   codeCoun,
		Address:    addressDetail,
		RoomNumber: roomNumber,
		IsDefault:  memberpb.AddressIsDefault(isDef),
		//Longitude:  longitude,
		//Latitude:   latitude,
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
	addressId, _ := strconv.ParseUint(m.PostForm("address_id"), 10, 64)
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)
	name := m.PostForm("name")
	mobile := m.PostForm("mobile")
	codeProv, _ := strconv.ParseUint(m.PostForm("code_prov"), 10, 64)
	codeCity, _ := strconv.ParseUint(m.PostForm("code_city"), 10, 64)
	codeCoun, _ := strconv.ParseUint(m.PostForm("code_coun"), 10, 64)
	addressDetail := m.PostForm("address")
	roomNumber := m.PostForm("room_number")
	isDefault := m.PostForm("is_default")
	//longitude := m.PostForm("longitude")
	//latitude := m.PostForm("latitude")
	
	isDef := 0
	if isDefault == "true" {
		isDef = 1
	}
	req := &memberpb.Address{
		AddressId:  addressId,
		MemberId:   memberId,
		Name:       name,
		Mobile:     mobile,
		CodeProv:   codeProv,
		CodeCity:   codeCity,
		CodeCoun:   codeCoun,
		Address:    addressDetail,
		RoomNumber: roomNumber,
		IsDefault:  memberpb.AddressIsDefault(isDef),
		//Longitude:  longitude,
		//Latitude:   latitude,
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
