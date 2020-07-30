
## 用户查询列表

#### 接口URL
> {{baseUrl}}/users

#### 请求方式
> GET

#### Content-Type
> multipart/form-data





#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | application/json |  选填 | - |

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
| Content-Type     | application/json |  选填 | - |

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
| Content-Type     | multipart/form-data |  选填 | - |

#### 请求Body参数

```javascript
{
    "user_name": "heyong",
    "password": "000000"
}
```



## 用户退出

#### 接口URL
> {{baseUrl}}/users/session

#### 请求方式
> DELETE

#### Content-Type
> multipart/form-data









## 修改密码

#### 接口URL
> {{baseUrl}}/users/session/password

#### 请求方式
> PUT

#### Content-Type
> multipart/form-data





#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| Content-Type     | multipart/form-data |  选填 | - |

#### 请求Body参数

```javascript
{
    "password": "000000"
}
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
| Content-Type     | application/json |  选填 | - |

#### 请求Body参数

```javascript
{
    "real_name": "陈凡2121",
    "role_code": ["admin"],
    "user_name": "chengfa2n1"
}
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
| userId     | d2518c70-fa26-4a9d-bf77-ddc60651d3e2 | - |






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
| userId     | d2518c70-fa26-4a9d-bf77-ddc60651d3e2 | - |






## 重置密码

#### 接口URL
> {{baseUrl}}/users/:userId/password

#### 请求方式
> DELETE

#### Content-Type
> multipart/form-data



#### 路径参数

| 参数        | 示例值   |  参数描述  |
| :--------   | :-----  | :----  |
| userId     | d2518c70-fa26-4a9d-bf77-ddc60651d3e2 | - |





