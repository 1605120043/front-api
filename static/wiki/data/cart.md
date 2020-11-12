购物车
-----------

### 1. <a id="add">加入购物车</a>

#### 接口功能

> 加入购物车，返回购物车列表

#### URL

> cart/add

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|product_id  |ture |int|商品id |
|product_spec_id  |true |int |规格id|
|nums  |true |int |加入数量|
|is_select  |true |bool |是否选中|
|is_plus  |false |bool |是否叠加，默认是叠加|


#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|amount | float | 商品金额 |
|promotion | float | 优惠金额 |
|cart_id | int | 购物车编号 |
|product_id | int | 商品编号 |
|product_spec_id | int | 规格编号 |
|stock | int | 库存 |
|product_name | string | 商品名称 |
|spec_name | string | 规格名称 |
|image | string | 规格图片 |
|price | string | 单个商品金额 |
|num | string | 购物车数量 |
|checked | bool | 是否选中 |
|product_amount | float | 单个商品总金额 ， price*num|

#### 接口示例
```
{
    "run_time": 0.821,
    "code": 1,
    "message": "",
    "data": [{
        "amount": 0,
        "promotion": 0,
        "products": [
            {
                "cart_id": 28,
                "product_id": 1,
                "product_spec_id": 11,
                "stock": 99,
                "product_name": "Hello",
                "spec_name": "99",
                "image": "sun54654sd6f4ds65f3ew",
                "price": 99,
                "num": 1,
                "checked": false,
                "product_amount": 99
            }
        ]
    }]
}
```


### 2. <a id="delete">删除购物车</a>

#### 接口功能

> 删除购物车，返回购物车列表

#### URL

> cart/delete

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|is_all  |ture |int|是否删除全部， 0或1 |
|cart_ids  |true |json string |删除的购物车id，如果is_all等于1，将忽略此字段数据, 例 : [40, 1, 2, 3]|


#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|amount | float | 商品金额 |
|promotion | float | 优惠金额 |
|cart_id | int | 购物车编号 |
|product_id | int | 商品编号 |
|product_spec_id | int | 规格编号 |
|stock | int | 库存 |
|product_name | string | 商品名称 |
|spec_name | string | 规格名称 |
|image | string | 规格图片 |
|price | string | 单个商品金额 |
|num | string | 购物车数量 |
|checked | bool | 是否选中 |
|product_amount | float | 单个商品总金额 ， price*num|

#### 接口示例
```
{
    "run_time": 0.113,
    "code": 1,
    "message": "",
    "data": [{
        "amount": 0,
        "promotion": 0,
        "products": []
    }]
}
```

### 3. <a id="index">购物车列表</a>

#### 接口功能

> 获取购物车列表

#### URL

> cart/index

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|amount | float | 商品金额 |
|promotion | float | 优惠金额 |
|cart_id | int | 购物车编号 |
|product_id | int | 商品编号 |
|product_spec_id | int | 规格编号 |
|stock | int | 库存 |
|product_name | string | 商品名称 |
|spec_name | string | 规格名称 |
|image | string | 规格图片 |
|price | string | 单个商品金额 |
|num | string | 购物车数量 |
|checked | bool | 是否选中 |
|product_amount | float | 单个商品总金额 ， price*num|

#### 接口示例
```
{
    "run_time": 0.821,
    "code": 1,
    "message": "",
    "data": [{
        "amount": 0,
        "promotion": 0,
        "products": [
            {
                "cart_id": 28,
                "product_id": 1,
                "product_spec_id": 11,
                "stock": 99,
                "product_name": "Hello",
                "spec_name": "99",
                "image": "sun54654sd6f4ds65f3ew",
                "price": 99,
                "num": 1,
                "checked": false,
                "product_amount": 99
            }
        ]
    }]
}
```

### 4. <a id="checked">选择购物车</a>

#### 接口功能

> 批量选择或取消购物车，返回购物车列表

#### URL

> cart/checked

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|cart_checked  |true |json string| [{"cart_id":8,"is_select":1}] |

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|amount | float | 商品金额 |
|promotion | float | 优惠金额 |
|cart_id | int | 购物车编号 |
|product_id | int | 商品编号 |
|product_spec_id | int | 规格编号 |
|stock | int | 库存 |
|product_name | string | 商品名称 |
|spec_name | string | 规格名称 |
|image | string | 规格图片 |
|price | string | 单个商品金额 |
|num | string | 购物车数量 |
|checked | bool | 是否选中 |
|product_amount | float | 单个商品总金额 ， price*num|

#### 接口示例
```
{
    "run_time": 0.821,
    "code": 1,
    "message": "",
    "data": [{
        "amount": 0,
        "promotion": 0,
        "products": [
            {
                "cart_id": 28,
                "product_id": 1,
                "product_spec_id": 11,
                "stock": 99,
                "product_name": "Hello",
                "spec_name": "99",
                "image": "sun54654sd6f4ds65f3ew",
                "price": 99,
                "num": 1,
                "checked": false,
                "product_amount": 99
            }
        ]
    }]
}
```


### 5. <a id="buy">立即购买</a>

#### 接口功能

> 立即购买，计算优惠

#### URL

> cart/buy

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|products  |true |json string| [{"cart_id":8,"product_id":1,"product_spec_id":1,"num":1}] |

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|pay_amount | float | 实付金额 |
|order_amount | float | 订单金额 |
|order_promotion | float | 订单优惠 |
|coupon_promotion | float | 优惠券优惠 |
|promotion_list | array | 优惠列表 |
|coupon_list | array | 优惠券列表 |
|order_weight | float | 订单总量 |
|cost_freight | float | 运费 |
|product_id | int | 商品id |
|product_spec_id | int | 规格id |
|stock | int | 库存 |
|product_name | string | 商品名称 |
|spec_name | string | 规格名称 |
|image | string | 商品规格图片 |
|price | float | 价格 |
|num | int | 数量 |
|weight | float | 重量 |
|product_weight | float | 总重量 |
|product_amount | float | 总价格 |
|product_promotion | float | 总优惠价格 |
|error | string | 错误信息 |

#### 接口示例
```
{
    "run_time": 0.373,
    "code": 1,
    "message": "",
    "data": [{
        "order_amount": 99,
        "order_promotion": 0,
        "coupon_promotion": 0,
        "promotion_list": [],
        "coupon_list": [],
        "order_weight": 99,
        "cost_freight": 0,
        "products": [
            {
                "product_id": 1,
                "product_spec_id": 11,
                "stock": 99,
                "product_name": "Hello",
                "spec_name": "99",
                "image": "sun54654sd6f4ds65f3ew",
                "price": 99,
                "num": 1,
                "weight": 99,
                "product_weight": 99,
                "product_amount": 99,
                "product_promotion": 0,
                "error": null
            }
        ]
    }]
}
```


### 6. <a id="count">获取购物车数量</a>

#### 接口功能

> 获取购物车数量

#### URL

> cart/count

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.001251953,
    "code": 1,
    "message": "",
    "data": [
        99
    ]
}
```