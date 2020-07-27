# 北冥有声小说 bm-novel

- [API接口]

## 功能

- 用户管理
- 权限管理
- 小说管理

## 组件
- db: postgres
- cache: redis
- auth: cookie + jwt + redis
- security: bcrypt
- dao: sqlx + entity + goqu
- web: chi + httpkit
- api: md + postman
- error: pkg/errors
- log: -

- 代码检查 
```shell script
revive -config=revive.toml -formatter friendly bm-novel/...
```

## 目录结构

```
├─cmd
│  └─server
├─configs
│  └─server
├─docs
└─internal
    ├─config        // 配置读取
    ├─controller    // web api
    │  └─user
    ├─domain        // 业务逻辑
    │  ├─notice
    │  ├─novel
    │  ├─permission
    │  └─user
    ├─http          // http 组件
    │  ├─auth       // 身份认证与授权验证
    │  └─web        // web 组件
    └─infrastructure    // 基础设施
        ├─cookie        
        ├─persistence   // 数据库持久化
        │  ├─permission
        │  └─user
        ├─redis         // 缓存
        └─security      // 安全/加密

```

[API接口]: https://gitlab.haochang.tv/heyong/bm-novel/issues?scope=all&utf8=%E2%9C%93&state=opened&label_name[]=documentation