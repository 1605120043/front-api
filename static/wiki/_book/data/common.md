公共
-----------

### 1. <a id="get-area-list">省市区</a>

#### 接口功能

> 获取省市区列表，已归类

#### URL

> common/get-area-list

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|label | string |名称 |
|value | int | 编号 |

#### 接口示例
```
{
    "run_time": 0.111, 
    "code": 1, 
    "message": "", 
    "data": [{
        "areas": [
            {
                "label": "上海市", 
                "value": 310000, 
                "children": [
                    {
                        "label": "上海市", 
                        "value": 310100, 
                        "children": [
                            {
                                "label": "杨浦区", 
                                "value": 310110
                            }, 
                            {
                                "label": "青浦区", 
                                "value": 310118
                            }, 
                            {
                                "label": "崇明区", 
                                "value": 310151
                            }, 
                            {
                                "label": "金山区", 
                                "value": 310116
                            }
                        ]
                    }
                ]
            }
        ]
    }]
}
```

### 2. <a id="mobile-login">手机号码登录</a>

#### 接口功能

> 会员根据手机号码登录

#### URL

> common/mobile-login

#### HTTP请求方式

> POST

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|mobile  |ture |string|手机号码 |
|code  |true |string |验证码, 当前暂无, 固定传“0000”|

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|token | string |登录token |
|expire | int | 过期时间 |
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
    "run_time": 0.141,
    "code": 1,
    "message": "",
    "data": [{
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNiwibW9iaWxlIjoiMTg2MjE1MjA2MDUiLCJleHAiOjE2MDMzMzgzNTYsImlhdCI6MTYwMzI1MTk1NiwiaXNzIjoiMTg2MjE1MjA2MDUifQ.mGTR_FH8CBmJmkexx4UHE9c9cY1FfHrEDjDIRr5Yf88",
        "expire": 1603338356,
        "info": {
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
        }
    }]
}
```

### 3. <a id="mobile-code">发送验证码</a>

#### 接口功能

> 发送验证码

#### URL

> common/send-code

#### HTTP请求方式

> POST

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|mobile  |ture |string|手机号码 |

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.005,
    "code": 1,
    "message": "",
    "data": []
}
```