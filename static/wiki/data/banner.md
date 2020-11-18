轮播图、广告位
-----------

### 1. <a id="index">轮播图、广告位</a>

#### 接口功能

> 获取轮播、广告位列表

#### URL

> banner/index

#### HTTP请求方式

> GET

#### 请求参数

无

#### 返回字段
|返回字段|字段类型|说明 |
|:----- |:------|:----------------------------- |
|banner | array |轮播图数据 |
|--id | int | 编号 |
|--ele_info | json |轮播信息 |
|---image_url | string |图片地址 |
|---redirect_url | string |跳转地址 |
|---sort | int |排序 |
|--tag_name | string |标识名 |
|ad | array |广告位数据 |
|--id | int |编号 |
|--ele_info | json |广告位信息 |
|---image_url | string |图片地址 |
|---redirect_url | string |跳转地址 |
|---sort | int |排序 |
|--tag_name | string |标识名 |

#### 接口示例
```json
{
	"run_time": 0.0140086,
	"code": 1,
	"message": "",
	"data": [{
		"banner": [{
			"id": 2,
            "ele_info": "[{\"image_url\":\"fasdfsd.jpg\",\"redirect_url\":\"http://www.baidu.com\",\"sort\":0}]",
            "tag_name": "home"
		}],
		"ad": [{
			"id": 3,
            "ele_info": "[{\"image_url\":\"fasdfsd.jpg\",\"redirect_url\":\"http://www.baidu.com\",\"sort\":0}]",
            "tag_name": "home"
		}]
	}]
}
```