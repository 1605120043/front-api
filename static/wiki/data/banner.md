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
	"run_time": 0.0089998,
	"code": 1,
	"message": "",
	"data": [{
		"banner": [{
			"id": 5,
			"ele_info": [{
				"image_url": "62e3943ecba913333d5144b839ff5e64.png",
				"redirect_url": "www.geng.com",
				"sort": 2
			}, {
				"image_url": "47dcff59681d3597d6da3ea5911efc3b.jpeg",
				"redirect_url": "www.baidu.com",
				"sort": 1
			}],
			"tag_name": "dog2"
		}, {
			"id": 2,
			"ele_info": [{
				"image_url": "62e3943ecba913333d5144b839ff5e64.png",
				"redirect_url": "www.geng.com",
				"sort": 2
			}, {
				"image_url": "47dcff59681d3597d6da3ea5911efc3b.jpeg",
				"redirect_url": "www.baidu.com",
				"sort": 1
			}],
			"tag_name": "home"
		}],
		"ad": [{
			"id": 6,
			"ele_info": [{
				"image_url": "62e3943ecba913333d5144b839ff5e64.png",
				"redirect_url": "www.geng.com",
				"sort": 2
			}],
			"tag_name": "dddd2"
		}, {
			"id": 3,
			"ele_info": [{
				"image_url": "62e3943ecba913333d5144b839ff5e64.png",
				"redirect_url": "www.geng.com",
				"sort": 2
			}],
			"tag_name": "RRRd"
		}]
	}]
}
```