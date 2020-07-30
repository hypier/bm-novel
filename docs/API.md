# 北冥有声API接口

## 用户管理

### 用户查询列表

#### URI

`GET /users`

- 可按角色、姓名搜索
- 分页查询

#### Request
``` json
{
    "role_code": (string),       // 角色代码
    "real_name": (string),       // 姓名
    "page_index": (integer),     // 当前页码
    "page_size": (integer)       // 每页数量
}
```

#### Response
```
[
    {
      "user_id": (uuid),         // 账号ID
      "user_name": (string),     // 账号
      "role_code": (string),     // 角色代码
      "real_name": (string),     // 姓名
      "lock": （boolean)        // 是否锁定
    }
]
```

- `200` 操作成功
- `401` 未登陆或超时
- `403` 没有权限操作

----------

### 创建用户

#### URI
`POST /users` 

#### Request
```json
{
  "role_code": (string),     // 角色代码
  "real_name": (string),     // 姓名
  "user_name": (string)      // 账号名
}
```

#### Response
- `201` 用户创建成功
- `401` 未登陆或超时
- `403` 没有权限操作
- `409` 用户名已存在

----------

### 编辑用户

#### URI
`PATCH /users/{user_id}`

#### Request
``` json
{
  "real_name": (string),    // 姓名
  "role_code": (string),    // 角色代码
  "user_name": (string)     // 账号名
}
```

#### Response
- `200` 编辑成功
- `401` 未登陆或超时
- `403` 没有权限操作

----------

### 用户锁定

#### URI
`POST /users/{user_id}/lock`

#### Request
无

#### Response
- `200` 锁定成功
- `401` 未登陆或超时
- `403` 没有权限操作
- `406` 状态已锁定，不能操作

----------

### 用户解锁

#### URI
`DELETE /users/{user_id}/lock`

#### Request
无

#### Response
- `200` 解锁成功
- `401` 未登陆或超时
- `403` 没有权限操作
- `406` 状态已解锁，不能操作

----------

### 重置密码

密码将重置为：**123456**

#### URI
`DELETE /users/{user_id}/password`

#### Reqeust
无

#### Response
- `200` 操作成功
- `401` 未登陆或超时
- `403` 没有权限操作

----------

### 用户登陆

#### URI
`POST /users/session`

#### Request
```json
{
    "user_name": (string),   //账号名
    "password": (string)    //密码
}
```

#### Response
- `200`     登陆成功，同时下发cookie
``` json
{
    "need_change_password": (boolean),   // 是否需要修改密码
    "real_name": (string),       // 姓名
    "user_id": (uuid),           // 账号ID
    "user_name": (string)        // 账号名
}

```

- `401` 用户名或密码错误
- `423` 用户已锁定


----------
### 登陆后重设密码

- 用于用户创建或重置密码后的首次登陆重新设置密码
- 重复密码由前端验证是否相等
- 需要判断登陆接口返回的 **need_change_password**值

#### URI
`PUT /users/session/password`

#### Reqeust
```json
{
    "password": (string)    // 密码
}
```

#### Response
- `200` 操作成功
- `401` 未登陆或超时
- `403` 没有权限操作

----------
### 用户注销

#### URI
`DELETE /users/session`

#### Request
无

#### Response
- `200` 注销成功
- `401` 未登陆或超时
- `403` 没有权限操作