package cart

type CartProducts struct {
	CartId        uint64  `json:"cart_id"`
	ProductId     uint64  `json:"product_id"`
	ProductSpecId uint64  `json:"product_spec_id"`
	Stock         uint64  `json:"stock"` // 库存
	ProductName   string  `json:"product_name"`
	SpecName      string  `json:"spec_name"`
	Image         string  `json:"image"`
	Price         float64 `json:"price"`
	Num           uint64  `json:"num"` //购买数量
	Checked       bool    `json:"checked"`
	ProductAmount float64 `json:"product_amount"` //商品总金额
}

type Carts struct {
	Count     uint64          `json:"count"`     //购物车数量
	Amount    float64         `json:"amount"`    //总金额
	Promotion float64         `json:"promotion"` //促销金额
	Products  []*CartProducts `json:"products"`
}

type BuyReq struct {
	CartId        uint64 `json:"cart_id"`
	ProductId     uint64 `json:"product_id"`
	ProductSpecId uint64 `json:"product_spec_id"`
	Num           uint64 `json:"num"`
}

type BuyProducts struct {
	ProductId        uint64  `json:"product_id"`
	ProductSpecId    uint64  `json:"product_spec_id"`
	Stock            uint64  `json:"stock"` // 库存
	ProductName      string  `json:"product_name"`
	SpecName         string  `json:"spec_name"`
	Image            string  `json:"image"`
	Price            float64 `json:"price"`
	Num              uint64  `json:"num"` //购买数量
	Weight           float64 `json:"weight"`
	ProductWeight    float64 `json:"product_weight"`    //商品总重
	ProductAmount    float64 `json:"product_amount"`    //商品总金额
	ProductPromotion float64 `json:"product_promotion"` //商品促销金额
	Error            error   `json:"error"`             //商品错误信息
}

type BuyRes struct {
	OrderAmount     float64        `json:"order_amount"`     //订单总金额
	OrderPromotion  float64        `json:"order_promotion"`  //订单促销金额
	CouponPromotion float64        `json:"coupon_promotion"` //优惠券金额
	PromotionList   []string       `json:"promotion_list"`   //促销列表
	CouponList      []string       `json:"coupon_list"`      //优惠券列
	OrderWeight     float64        `json:"order_weight"`     // 订单总重
	CostFreight     float64        `json:"cost_freight"`     //运费
	PayAmount       float64        `json:"pay_amount"`       // 实付
	Products        []*BuyProducts `json:"products"`
}
