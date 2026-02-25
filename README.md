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

#Front(Everything bug)
https://github.com/Amadeus-Crs/Seronium