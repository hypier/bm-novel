
## 用户登陆

#### 接口URL
> {{baseUrl}}/users/session

#### 请求方式
> POST

#### Content-Type
> multipart/form-data





#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | application/json |  必填 | - |

#### 请求Body参数

```javascript
{
	"user_name": "heyong",
	"password": "123456"
}
```
| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| user_name     | heyong |  必填 | 用户名 |
| password     | 123456 |  必填 | 密码 |

#### 成功响应示例
```javascript
200     登陆成功，同时下发cookie

{
	"user_id": "b283c07f-962d-4256-9f2e-a0c054d6ba01",
	"user_name": "heyong",
	"real_name": "陈凡",
	"need_change_password": true
}
```

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| user_id     | b283c07f-962d-4256-9f2e-a0c054d6ba01 | 用户ID |
| user_name     | heyong | 用户名 |
| real_name     | 陈凡 | 姓名 |
| need_change_password     | true | 是否需要修改密码 |

#### 错误响应示例
```javascript
401     未登陆或超时
403     没有权限操作
409     用户名已存在
```


## 创建用户

#### 接口URL
> {{baseUrl}}/users

#### 请求方式
> POST

#### Content-Type
> multipart/form-data





#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | application/json |  必填 | - |

#### 请求Body参数

```javascript
{
	"real_name": "陈凡",
	"role_code": [
		"admin",
		"www"
	],
	"user_name": "heyong"
}
```
| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| real_name     | 陈凡 |  必填 | 姓名 |
| role_code     | ["admin","www"] |  必填 | 角色代码 |
| user_name     | heyong |  必填 | 用户名 |

#### 成功响应示例
```javascript
201 用户创建成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
409 用户名已存在
```


## 用户查询列表
- 参数实际为 body
#### 接口URL
> {{baseUrl}}/users?role_code=["admin"]&real_name=&page_index=1&page_size=2

#### 请求方式
> GET

#### Content-Type
> multipart/form-data

#### 请求Query参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| role_code     | ["admin"] | 选填 | - |
| real_name     | - | 选填 | - |
| page_index     | 1 | 必填 | - |
| page_size     | 2 | 必填 | - |




#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | application/json |  必填 | - |

#### 请求Body参数

```javascript
{
	"role_code": [
		"admin"
	],
	"page_index": 1,
	"page_size": 3
}
```

#### 成功响应示例
```javascript
[
	{
		"user_id": "082d7b86-0b5d-4475-8955-61b54b67f375",
		"user_name": "heyong5121122",
		"role_code": [
			"admin"
		],
		"real_name": "陈凡",
		"lock": false
	}
]
```

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| user_id     | 082d7b86-0b5d-4475-8955-61b54b67f375 | 用户ID |
| user_name     | heyong5121122 | 用户名 |
| role_code     | admin | 角色代码 |
| real_name     | 陈凡 | 姓名 |
| lock     | false | 是否锁定 |

#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
```


## 编辑用户

#### 接口URL
> {{baseUrl}}/users/:userId

#### 请求方式
> PATCH

#### Content-Type
> multipart/form-data



#### 路径参数

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| userId     | c67ee30c-9c76-4a2b-8d6c-063f114c616c | - |


#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | application/json |  必填 | - |

#### 请求Body参数

```javascript
{
	"real_name": "陈凡2121",
	"role_code": [
		"admin"
	],
	"user_name": "chengfa2n1"
}
```
| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| real_name     | 陈凡2121 |  必填 | 姓名 |
| role_code     | ["admin","www"] |  必填 | 角色代码 |
| user_name     | chengfa2n1 |  必填 | 用户名 |

#### 成功响应示例
```javascript
200 编辑成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
404 用户不存在
```


## 登陆后重设密码
* 用于用户创建或重置密码后的首次登陆重新设置密码
* 重复密码由前端验证是否相等
* 需要判断登陆接口返回的 need_change_password值
#### 接口URL
> {{baseUrl}}/users/session/password

#### 请求方式
> PUT

#### Content-Type
> multipart/form-data





#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | multipart/form-data |  必填 | - |

#### 请求Body参数

```javascript
{
	"password": "000000"
}
```
| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| password     | 000000 |  必填 | 密码 |

#### 成功响应示例
```javascript
200 操作成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
404 用户不存在
```


## 重置密码
* 密码将重置为：123456
#### 接口URL
> {{baseUrl}}/users/:userId/password

#### 请求方式
> DELETE

#### Content-Type
> multipart/form-data



#### 路径参数

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| userId     | b283c07f-962d-4256-9f2e-a0c054d6ba01 | 用户ID |



#### 请求Body参数

```javascript

```

#### 成功响应示例
```javascript
200 操作成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
404 用户不存在
```


## 用户锁定

#### 接口URL
> {{baseUrl}}/users/:userId/lock

#### 请求方式
> POST

#### Content-Type
> multipart/form-data



#### 路径参数

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| userId     | b283c07f-962d-4256-9f2e-a0c054d6ba01 | - |




#### 成功响应示例
```javascript
200 锁定成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
404 用户不存在
406 状态已锁定，不能操作
```


## 用户解锁

#### 接口URL
> {{baseUrl}}/users/:userId/lock

#### 请求方式
> DELETE

#### Content-Type
> multipart/form-data



#### 路径参数

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| userId     | b283c07f-962d-4256-9f2e-a0c054d6ba01 | 用户ID |




#### 成功响应示例
```javascript
200 解锁成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
404 用户不存在
406 状态已解锁，不能操作
```


## 用户注销

#### 接口URL
> {{baseUrl}}/users/session

#### 请求方式
> DELETE

#### Content-Type
> multipart/form-data







#### 成功响应示例
```javascript
200 注销成功
```


#### 错误响应示例
```javascript
401 未登陆或超时
403 没有权限操作
```

