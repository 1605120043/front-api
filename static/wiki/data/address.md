收货地址
-----------

### 1. <a id="index">收货地址</a>

#### 接口功能

> 根据登录会员获取收货地址，获取用户默认地址page_size可以传1

#### URL

> address/index

#### HTTP请求方式

> GET

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|page_size  |false |int|条数, 获取用户默认地址page_size可以传1 |

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|address_id | int |编号 |
|name | string | 收货人 |
|mobile | string |收货人手机号码 |
|address | string |收货详细地址 |
|is_default | bool |是否默认 |

#### 接口示例
```
{
    "run_time": 0.159,
    "code": 1,
    "message": "",
    "data": [
        {
            "address_id": 5,
            "name": "shrimp",
            "mobile": "222",
            "address": "上海市上海市长宁区中山国际广场B座2楼",
            "is_default": true
        },
        {
            "address_id": 4,
            "name": "1",
            "mobile": "1111",
            "address": "3333333",
            "is_default": true
        }
    ]
}
```

### 2. <a id="detail">收货地址详情</a>

#### 接口功能

> 根据收货地址id获取收货地址详情

#### URL

> address/detail

#### HTTP请求方式

> GET

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|address_id  |true |int|收货地址id |

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|address_id | int |编号 |
|name | string | 收货人 |
|mobile | string |收货人手机号码 |
|code_prov | int |省code |
|code_city | int |市code |
|code_coun | int |区code |
|code_prov_name | string |省名称 |
|code_city_name | string |市名称 |
|code_coun_name | string |区名称 |
|address | string |收货地址 |
|room_number | string |门牌号或楼号 |
|is_default | bool |是否默认 |

#### 接口示例
```
{
    "run_time": 0.123,
    "code": 1,
    "message": "",
    "data": [{
        "address_id": 4,
        "name": "1",
        "mobile": "11111",
        "code_prov": 2,
        "code_city": 2,
        "code_coun": 2,
        "code_prov_name": "",
        "code_city_name": "",
        "code_coun_name": "",
        "address": "33",
        "room_number": "33333",
        "is_default": true
    }]
}
```

### 3. <a id="add">添加收货地址</a>

#### 接口功能

> 添加收货地址

#### URL

> address/add

#### HTTP请求方式

> POST

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|name  |true |string|收货人 |
|mobile  |true |string|收货人手机号码 |
|code_prov  |true |int|省code |
|code_city  |true |int|市code |
|code_coun  |true |int|区code |
|address  |true |string|收货地址 |
|room_number  |true |string|门牌号或楼号 |
|is_default  |true |bool|区code |

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.123,
    "code": 1,
    "message": "",
    "data": []
}
```

### 4. <a id="edit">编辑收货地址</a>

#### 接口功能

> 编辑收货地址

#### URL

> address/edit

#### HTTP请求方式

> POST

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|address_id  |true |int|收货地址id |
|name  |true |string|收货人 |
|mobile  |true |string|收货人手机号码 |
|code_prov  |true |int|省code |
|code_city  |true |int|市code |
|code_coun  |true |int|区code |
|address  |true |string|收货地址 |
|room_number  |true |string|门牌号或楼号 |
|is_default  |true |bool|区code |

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.123,
    "code": 1,
    "message": "",
    "data": []
}
```

### 5. <a id="delete">删除收货地址</a>

#### 接口功能

> 删除收货地址

#### URL

> address/delete

#### HTTP请求方式

> POST

#### 请求参数

|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|address_id  |true |int|收货地址id |

#### 返回字段

#### 接口示例
```
{
    "run_time": 0.123,
    "code": 1,
    "message": "",
    "data": []
}
```