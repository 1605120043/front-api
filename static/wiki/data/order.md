订单
-----------

### 1. <a id="index">订单列表</a>

#### 接口功能

> 根据登录会员获取订单列表

#### URL

> order/index

#### HTTP请求方式

> GET

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|page  |false |int|页码 |
|page_size  |false |int|条数 |
|order_status  |false |int|订单状态( 全部（0）  待付款（1）  待发货（3）   待收货（4）) |


#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|order_id | int |订单id |
|grand_total | float |合计(应收总金额) |
|order_status | int |订单状态 1,待付款 2,待审核 3,待发货 4,待收货 5,完成 6,待评价 7,取消 |
|order_status_name | string |订单状态名称 |
|order_item_id | int |订单明细编号 |
|product_id | int |商品id |
|name | string |商品名称 |
|sku | string |SKU名称 |
|image | string |规格图片 |
|price | float |销售价 |
|old_price | float |原始价 |
|total_payable | float |应付总金额 |
|total_discount_amount | float |折扣总金额 |
|qty_ordered | int |订购数量 |
|weight | float |重量 |
|volume | float |体积 |
|spec | json string |规格详情 |
|qty_shipped | int |已发数量 |

#### 接口示例
```
{
    "run_time": 0.038,
    "code": 1,
    "message": "",
    "data": [{
        "total": 14,
        "orders": [
            {
                "order_id": 116680644523855873,
                "grand_total": 5354,
                "total_qty_ordered": 0,
                "order_status": 1,
                "order_status_name": "待付款",
                "order_items": [
                    {
                        "order_item_id": 116680644729376773,
                        "product_id": 30,
                        "name": "问问",
                        "sku": "453",
                        "image": "6dc289dfcf3e26cbe793c6d69ab533e8-w300.jpg",
                        "price": 5354,
                        "old_price": 453,
                        "total_payable": 5354,
                        "total_discount_amount": 0,
                        "qty_ordered": 1,
                        "weight": 453,
                        "volume": 453,
                        "spec": "[{\"name\":\"\",\"spec_value_id\":0,\"value\":\"\"}]",
                        "qty_shipped": 0
                    }
                ],
                "created_at": "2020-10-16 17:59:55"
            }
        ]
    }]
}
```

### 2. <a id="info">订单详情</a>

#### 接口功能

> 根据订单id获取订单详情

#### URL

> order/info

#### HTTP请求方式

> GET

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|order_id  |ture |int|订单id |

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|order_id | int |订单id |
|grand_total | float |合计(应收总金额) |
|subtotal | float |小计(商品总金额，未扣除折扣) |
|grand_total | float |合计(应收总金额) |
|total_paid | float |总付款金额 |
|shipping_amount | float |运费 |
|discount_amount | float |折扣 |
|payment_type | int |付款类型 1,在线支付 |
|payment_status | int |付款状态 1,未付款 2,已付款 3,部分付款 |
|payment_time | string |付款时间 |
|shipping_status | int |发货状态 1,未发货 2,部分发货 3,已发货 |
|shipping_time | string |发货时间 |
|confirm | int |是否确认收货 1,否 2,是 |
|confirm_time | string |确认收货时间 |
|order_status | int |订单状态 1,待付款 2,待审核 3,待发货 4,待收货 5,完成 6,待评价 7,取消 |
|refund_status | int |退款状态 1,未退款 2,部分退款 3,全部退款 |
|return_status | int |退货状态 1,未退货 2,部分退货 3,全部退货 |
|user_note | string |用户备注 |
|order_item_id | int |订单明细编号 |
|product_id | int |商品id |
|name | string |商品名称 |
|sku | string |SKU名称 |
|image | string |规格图片 |
|price | float |销售价 |
|old_price | float |原始价 |
|total_payable | float |应付总金额 |
|total_discount_amount | float |折扣总金额 |
|qty_ordered | int |订购数量 |
|weight | float |重量 |
|volume | float |体积 |
|spec | json string |规格详情 |
|qty_shipped | int |已发数量 |
|receiver | string |收货人 |
|telephone | string |收货人电话 |
|province | string |省 |
|city | string |市 |
|region | string |区 |
|street | string |街道 |
|order_shipment | string |订单快递信息 |
|created_at | string |订单生成时间 |

#### 接口示例
```
{
    "run_time": 0.048,
    "code": 1,
    "message": "",
    "data": [{
        "order_id": 116680644523855873,
        "subtotal": 5354,
        "grand_total": 5354,
        "total_paid": 0,
        "shipping_amount": 0,
        "discount_amount": 0,
        "payment_type": 1,
        "payment_status": 1,
        "payment_time": "",
        "shipping_status": 1,
        "shipping_time": "",
        "confirm": 1,
        "confirm_time": "",
        "order_status": 1,
        "refund_status": 1,
        "return_status": 1,
        "user_note": "Hello",
        "order_items": [
            {
                "order_item_id": 116680644729376773,
                "product_id": 30,
                "name": "问问",
                "sku": "453",
                "image": "6dc289dfcf3e26cbe793c6d69ab533e8-w300.jpg",
                "price": 5354,
                "old_price": 453,
                "total_payable": 5354,
                "total_discount_amount": 0,
                "qty_ordered": 1,
                "weight": 453,
                "volume": 453,
                "spec": "[{\"name\":\"\",\"spec_value_id\":0,\"value\":\"\"}]",
                "qty_shipped": 0
            }
        ],
        "order_address": {
            "order_address_id": 116680644830040071,
            "receiver": "shrimp23",
            "telephone": "18621520605",
            "province": "上海市",
            "city": "上海市",
            "region": "长宁区",
            "street": "123456123456"
        },
        "order_shipment": null,
        "created_at": "2020-10-16 17:59:55"
    }]
}
```

### 3. <a id="get-user-order-status-count">订单统计</a>

#### 接口功能

> 统计订单状态

#### URL

> order/get-user-order-status-count

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|order_status | int |订单状态 1,待付款 2,待审核 3,待发货 4,待收货 5,完成 6,待评价 7,取消 |
|count | int |数量 |

#### 接口示例
```
{
    "run_time": 0.0070006,
    "code": 1,
    "message": "",
    "data": [
        {
            "order_status": 1,
            "count": 1
        }
    ]
}
```

### 4. <a id="create-order">创建订单</a>

#### 接口功能

> 创建订单

#### URL

> order/create-order

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|products  |true |json string| [{"cart_id":8,"product_id":1,"product_spec_id":1,"num":1}] |
|address_id  |true |int|收货地址id |
|node  |false |string|备注 |

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|id | int |订单id |
|state | int |状态 |

#### 接口示例
```
{
    "run_time": 0.191, 
    "code": 1, 
    "message": "", 
    "data": [{
        "id": 201021143805141, 
        "state": 1
    }]
}
```

### 5. <a id="cancel-order">取消订单</a>

#### 接口功能

> 取消订单

#### URL

> order/cancel-order

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|order_id  |ture |int|订单id |

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.016669144,
    "code": 1,
    "message": "",
    "data": []
}
```

### 5. <a id="delete-order">删除订单</a>

#### 接口功能

> 删除订单

#### URL

> order/delete-order

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|order_id  |ture |int|订单id |

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.016669144,
    "code": 1,
    "message": "",
    "data": []
}
```