package order

import "github.com/shinmigo/pb/orderpb"

type AddressOrderRes struct {
	Receiver  string `json:"receiver"`
	Telephone string `json:"telephone"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Street    string `json:"street"`
}

type PaymentOrderRes struct {
	PaymentCode string `json:"payment_code"`
	PaymentName string `json:"payment_name"`
}

type ShipmentOrderRes struct {
	CarrierName    string `json:"carrier_name"`
	TrackingNumber string `json:"tracking_number"`
}

type ListDetailOrderRes struct {
	OrderId         uint64                            `json:"order_id"`
	GrandTotal      float64                           `json:"grand_total"`
	TotalQtyOrdered uint64                            `json:"total_qty_ordered"`
	OrderStatus     orderpb.OrderStatus               `json:"order_status"`
	OrderStatusName string                            `json:"order_status_name"`
	OrderItems      []*orderpb.OrderDetail_OrderItems `json:"order_items"`
	CreatedAt       string                            `json:"created_at"`
}

type DetailOrderRes struct {
	OrderId        uint64                             `json:"order_id"`
	Subtotal       float64                            `json:"subtotal"`
	GrandTotal     float64                            `json:"grand_total"`
	TotalPaid      float64                            `json:"total_paid"`
	ShippingAmount float64                            `json:"shipping_amount"`
	DiscountAmount float64                            `json:"discount_amount"`
	PaymentType    orderpb.OrderPaymentType           `json:"payment_type"`
	PaymentStatus  orderpb.OrderPaymentStatus         `json:"payment_status"`
	PaymentTime    string                             `json:"payment_time"`
	ShippingStatus orderpb.OrderShippingStatus        `json:"shipping_status"`
	ShippingTime   string                             `json:"shipping_time"`
	Confirm        orderpb.OrderConfirm               `json:"confirm"`
	ConfirmTime    string                             `json:"confirm_time"`
	OrderStatus    orderpb.OrderStatus                `json:"order_status"`
	RefundStatus   orderpb.OrderRefundStatus          `json:"refund_status"`
	ReturnStatus   orderpb.OrderReturnStatus          `json:"return_status"`
	UserNote       string                             `json:"user_note"`
	OrderItems     []*orderpb.OrderDetail_OrderItems  `json:"order_items"`
	OrderAddress   *orderpb.OrderDetail_OrderAddress  `json:"order_address"`
	OrderShipment  *orderpb.OrderDetail_OrderShipment `json:"order_shipment"`
	CreatedAt      string                             `json:"created_at"`
}

type ListOrderRes struct {
	Total  uint64                `json:"total"`
	Orders []*ListDetailOrderRes `json:"orders"`
}

type UserOrderStatusCountRes struct {
	OrderStatus uint64 `json:"order_status"`
	Count       uint64 `json:"count"`
}
