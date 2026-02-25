# How To Use
```bash
cd #项目根目录
docker-compose up -d
go run ./cmd/server/main.go
```

# Some Basic Test
```JSON
//authgroup
//POST/api/auth/register
{
    "username":"",
    "password":"",
}

//POST/api/auth/login
{
    "username":"",
    "password":"",
}
```
以下默认Authorization:Bearer {{Token}}
```JSON
//apigroup
//GET/api/user/profile

//POST/api/posts
{
    "title":"",
    "content":"",
    "type":"",//question or article
}

//GET/api/posts/:id

//GET/api/posts
//Query
//sort = hot //or orthers
//offset = 0
//limit = 10

//POST/api/likes
{
    "target_type": "",
    "target_id": //int
}

//POST/api/comments
{
  "target_type": "",//post or answer
  "target_id": ,//int
  "content": ""
}

//PUT/api/posts/:id
{
  "title": "",
  "content": ""
}

//DELETE/api/posts/:id

//POST/api/collections
{
  "post_id": //int
}

//POST/api/upload
//form-data
//参数名'file' , 参数值 image only
```

# Front(Everything bug)
前端部分是自己写的，仍存在大量运行bug，锐意更新中

https://github.com/Amadeus-Crs/Seronium

# Feature
- 实现基本的用户注册与登录
- 基本的用户鉴权
- 密码的加盐加密
- 发布/删除/更新/获取文章、问题
- 文章可以使用markdown
- 能够对文章、问题进行回复/评论
- 可以对文章、问题回答进行点赞和收藏
- 按照问题热度进行排序
- 使用dockerfile文件打包
- 考虑SQL注入
- 通过viper或者其他方式加载相关信息的配置文件
- 通过zap集成日志