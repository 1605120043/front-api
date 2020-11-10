商品
-----------

### 1. <a id="tag">商品标签</a>

#### 接口功能

> 获取商品标签列表

#### URL

> product/tag

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|id | int | 编号 |
|name | string | 名称 |

#### 接口示例
```
{
    "run_time": 0.02,
    "code": 1,
    "message": "",
    "data": [
        {
            "id": 22,
            "name": "上海IE"
        },
        {
            "id": 18,
            "name": "新品"
        },
        {
            "id": 19,
            "name": "满二赠一"
        },
        {
            "id": 20,
            "name": "满五赠一"
        }
    ]
}
```

### 2. <a id="index">商品列表</a>

#### 接口功能

> 获取商品列表

#### URL

> product/index

#### HTTP请求方式

> GET

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|category_id  |false |int|分类id |
|tag_id  |false |int |标签id|
|name  |false |string |商品名称|
|page  |false |int |页码，默认1|
|page_size  |false |int |条数，默认20|


#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|product_id | int | 编号 |
|name | string | 名称 |
|image | string | 商品列表图片 |
|price | float | 商品价格，小数点保留2位 |
|short_description | string | 商品简介 |

#### 接口示例
```
{
    "run_time": 0.107,
    "code": 1,
    "message": "",
    "data": [
        {
            "product_id": 30,
            "name": "问问",
            "image": "d00cd077d140ac46152519692c59dd9e-w300.jpg",
            "price": 4,
            "short_description": "请我"
        },
        {
            "product_id": 13,
            "name": "二维",
            "image": "d00cd077d140ac46152519692c59dd9e-w300.jpg",
            "price": 2,
            "short_description": "请我"
        },
        ...
    ]
}
```

### 3. <a id="detail">商品详情</a>

#### 接口功能

> 根据商品id获取商品详情

#### URL

> product/detail

#### HTTP请求方式

> GET

#### 请求参数
|参数|必选|类型|说明|
|:----- |:-------|:-----|----- |
|product_id  |true |int|商品id |

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|product_id | int | 编号 |
|category_id | int | 分类id |
|kind_id | int | 类型id |
|short_description | string | 商品简介 |
|unit | string | 单位 |
|images | array | 商品图片 |
|spec_type | int | 规格类型 1, 单规格 2, 多规格 |
|status | int | 上下架状态：1上架，2下架 |
|tags | array | 标签 |
|spec | array | 商品所有规格 |
|param | array | 商品所有参数 |
|description | string | 商品描述 |
|category_name | string | 分类名称 |
|kind_name | string | 类型名称 |
|price | float | 商品价格，小数点保留2位 |
|spec_description | json string | 规格选择器 |
|param_description | json string | 参数选择器 |

#### 接口示例
```
{
    "run_time": 0.052,
    "code": 1,
    "message": "",
    "data": [{
        "product_id": 24,
        "category_id": 32,
        "kind_id": 16,
        "name": "性感短裙",
        "short_description": "性感短裙你值得拥有",
        "unit": "件",
        "images": [
            "d00cd077d140ac46152519692c59dd9e-w300.jpg",
            "9a03f957e0fa253ee2a632fa95c02fff-w300.jpg",
            "32bedb45662130484792a66d355e7ad9-w300.jpg"
        ],
        "spec_type": 2,
        "status": 1,
        "tags": [
            18
        ],
        "spec": [
            {
                "image": "cc60a023cf41f9bdcc6c74474c3ff217-w300.jpg",
                "price": 5,
                "old_price": 5,
                "cost_price": 5,
                "stock": 5,
                "sku": "5",
                "weight": 5,
                "volume": 5,
                "spec_value_id": [
                    55,
                    143,
                    150
                ],
                "spec_value_id_str": "55,143,150",
                "product_spec_id": 101
            },
            {
                "image": "cc60a023cf41f9bdcc6c74474c3ff217-w300.jpg",
                "price": 6,
                "old_price": 6,
                "cost_price": 6,
                "stock": 6,
                "sku": "6",
                "weight": 6,
                "volume": 6,
                "spec_value_id": [
                    56,
                    143,
                    150
                ],
                "spec_value_id_str": "56,143,150",
                "product_spec_id": 102
            },
            {
                "image": "cc60a023cf41f9bdcc6c74474c3ff217-w300.jpg",
                "price": 8,
                "old_price": 8,
                "cost_price": 8,
                "stock": 8,
                "sku": "8",
                "weight": 8,
                "volume": 8,
                "spec_value_id": [
                    56,
                    145,
                    150
                ],
                "spec_value_id_str": "56,145,150",
                "product_spec_id": 103
            }
        ],
        "param": [
            {
                "param_id": 7,
                "value": "52"
            },
            {
                "param_id": 9,
                "value": "60"
            },
            {
                "param_id": 10,
                "value": "62"
            }
        ],
        "description": "[{\"type\":2,\"text\":\"0529643c06a48c2f5651fae0d06216f8-w300.jpg\"},{\"type\":1,\"text\":\"性感短裙，魅力逼人\"},{\"type\":2,\"text\":\"0529643c06a48c2f5651fae0d06216f8-w300.jpg\"},{\"type\":1,\"text\":\"女装性感，美丽人生\"}]",
        "category_name": "栗子",
        "kind_name": "女装",
        "price": 5,
        "category_path": [
            "29",
            "33",
            "32"
        ],
        "spec_description": "[{\"name\": \"颜色\", \"spec_id\": 14, \"children\": {\"150\": {\"content\": \"白色\", \"spec_id\": 14, \"spec_value_id\": 150}}}, {\"name\": \"尺码\", \"spec_id\": 13, \"children\": {\"55\": {\"content\": \"M\", \"spec_id\": 13, \"spec_value_id\": 55}, \"56\": {\"content\": \"L\", \"spec_id\": 13, \"spec_value_id\": 56}}}, {\"name\": \"材质\", \"spec_id\": 11, \"children\": {\"143\": {\"content\": \"蕾丝\", \"spec_id\": 11, \"spec_value_id\": 143}, \"145\": {\"content\": \"纯棉\", \"spec_id\": 11, \"spec_value_id\": 145}}}]",
        "param_description": "[{\"name\": \"产地\", \"children\": {\"52\": {\"content\": \"上海\", \"param_id\": 7, \"param_value_id\": 52}}, \"param_id\": 7}, {\"name\": \"系列\", \"children\": {\"60\": {\"content\": \"美人系列\", \"param_id\": 9, \"param_value_id\": 60}}, \"param_id\": 9}, {\"name\": \"款式\", \"children\": {\"62\": {\"content\": \"短裙\", \"param_id\": 10, \"param_value_id\": 62}}, \"param_id\": 10}]"
    }]
}
```

### 4. <a id="category">分类</a>

#### 接口功能

> 获取商品分类列表

#### URL

> category/index

#### HTTP请求方式

> GET

#### 请求参数

#### 返回字段

|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|id | int | 编号 |
|pid | int | 父类编号 |
|name | string | 名称 |
|icon | string | 分类icon |
|sort | int | 排序值 |

#### 接口示例
```
{
    "run_time": 0.047,
    "code": 1,
    "message": "",
    "data": [
        {
            "id": 99,
            "pid": 98,
            "name": "酱油",
            "icon": "2F9a504fc2d56285352d7d584290ef76c6a7ef6330",
            "sort": 2
        },
        {
            "id": 98,
            "pid": 0,
            "name": "中山一米",
            "icon": "2F9a504fc2d56285352d7d584290ef76c6a7ef6330",
            "sort": 2
        }
    ]
}
```