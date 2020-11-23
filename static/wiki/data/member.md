会员
-----------

### 1. <a id="info">会员详情</a>

#### 接口功能

> 获取登录会员详情

#### URL

> member/info

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|member_id | int |用户id |
|nickname | string |用户昵称 |
|mobile | string |手机号码 |
|name | string |姓名 |
|gender | int |性别 0未知, 1女 2男 |
|id_card | string |身份证 |
|birthday | string |出生年月 |
|avatar | string |头像 |
|email | string |邮箱 |
|member_level_id | int |会员等级ID |
|point | int |积分 |
|balance | float |余额，小数点2位 |

#### 接口示例
```
{
    "run_time": 0.009,
    "code": 1,
    "message": "",
    "data": [{
        "member_id": 16,
        "nickname": "shrimp",
        "mobile": "18621520605",
        "name": "",
        "gender": 0,
        "id_card": "",
        "birthday": "0001-01-01T00:00:00Z",
        "avatar": "",
        "email": "",
        "member_level_id": 0,
        "point": 0,
        "balance": 0
    }]
}
```


### 2. <a id="update">编辑会员</a>

#### 接口功能

> 编辑登录会员详情

#### URL

> member/update

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|nickname  |true |string|用户昵称 |
|gender  |true |int |性别 0未知, 1女 2男|
|birthday  |false |string |出生年月 |
|avatar  |true |string |头像|

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

### 3. <a id="payment">支付方式</a>

#### 接口功能

> 获取支付方式列表

#### URL

> member/payment

#### HTTP请求方式

> POST

#### 请求参数

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|name | string |支付名称 |
|code | string |支付code |
|status | int |启用状态 1=启用 2=停用 |

#### 接口示例
```
{
    "run_time": 0,
    "code": 1,
    "message": "",
    "data": [
        {
            "name": "微信支付",
            "code": "Wechat",
            "status": 1
        },
        {
            "name": "支付宝支付",
            "code": "Alipay",
            "status": 1
        }
    ]
}
```

### 4. <a id="pay">支付</a>

#### 接口功能

> 支付接口

#### URL

> member/pay

#### HTTP请求方式

> POST

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|order_id  |true |string|订单id |
|payment_code  |true |string |支付方式|
|trade_type  |true |string |支付形式， 支付宝：WAP；...|

#### 微信返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|appid | string |小程序appid |
|partnerid | string |商户号id |
|prepayid | string |预支付交易会话标识 |
|noncestr | string |随机字符串 |
|package | string |签名字段 |
|timestamp | string |时间戳 |
|signType | string |加密方式 |
|paySign | string |签名 |


#### 接口示例
```

```