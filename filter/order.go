package filter

import (
	"goshop/front-api/model/order"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Order struct {
	validation validation.Validation
	*gin.Context
}

func NewOrder(c *gin.Context) *Order {
	return &Order{Context: c, validation: validation.Validation{}}
}

func (m *Order) Index() (list *order.ListOrderRes, err error) {
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	orderStatus := m.DefaultQuery("order_status", "0")
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)

	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("拉取格式不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("拉取格式不正确")
	valid.Match(orderStatus, regexp.MustCompile(`^[0-9]+$`)).Message("订单状态错误")
	if valid.HasError() {
		return nil, valid.GetError()
	}

	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	orderStatusNum, _ := strconv.ParseUint(orderStatus, 10, 64)
	listCategoryReq := &orderpb.ListOrderReq{
		Page:        pageNum,
		PageSize:    pageSizeNum,
		MemberId:    memberId,
		OrderStatus: orderpb.OrderStatus(orderStatusNum),
	}
	list, err = service.NewOrder(m.Context).Index(listCategoryReq)
	return
}

func (m *Order) Info() (*order.DetailOrderRes, error) {
	orderId := m.Query("order_id")
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)

	valid := validation.Validation{}
	valid.Required(orderId).Message("请选择订单！")
	valid.Match(orderId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("订单信息错误！")
	if valid.HasError() {
		return nil, valid.GetError()
	}

	orderIdNum, _ := strconv.ParseUint(orderId, 10, 64)
	return service.NewOrder(m.Context).Info(memberId, orderIdNum)
}

func (m *Order) GetUserOrderStatusCount() (list []*order.UserOrderStatusCountRes, err error) {
	memberId, _ := strconv.ParseUint(m.GetString("goshop_member_id"), 10, 64)

	list, err = service.NewOrder(m.Context).GetUserOrderStatusCount(memberId)
	return
}
