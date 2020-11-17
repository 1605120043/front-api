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
|--image_url | string |轮播图图片地址 |
|--redirect_url | string |轮播图跳转地址 |
|--sort | int |排序 |
|ad | array |广告位数据 |
|--id | int |编号 |
|--image_url | string |广告位图片地址 |
|--redirect_url | string |广告位跳转地址 |
|--sort | int |排序 |

#### 接口示例
```json
{
	"run_time": 0.0140086,
	"code": 1,
	"message": "",
	"data": [{
		"banner": [{
			"id": 2,
			"image_url": "sdfdsfdsa.jpg",
			"redirect_url": "http://fsdefasfsdf.jpg",
			"sort": 1
		}],
		"ad": [{
			"id": 3,
			"image_url": "sdfdsfdsfadfdsa.jpg",
			"redirect_url": "http://fsdefasfsdf.jpg",
			"sort": 1
		}]
	}]
}
```