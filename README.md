## DuckyGo MVC
   
   这是Django么？ 不是的！这是采用gin + pongo2gin + gorm搭建的MVC框架~

## 第三方模块导航
- [gin](https://github.com/gin-gonic/gi://github.com/gin-gonic/gin)
- [pongo2gin](https://gitlab.com/go-box/pongo2gin)
- [gin-sessions](https://github.com/gin-contrib/sessions)
- [go-redis](https://github.com/go-redis/redis)
- [gorm](https://github.com/jinzhu/gorm)
- [godotenv](https://github.com/joho/godotenv)

## 结构说明
- `conf` 用于服务器配置和url配置
- `controlls` 业务handler
- `models` 数据库建模
- `cache` 缓存模块
- `services` 服务 放置可重用的逻辑
- `middlewares` 中间件
- `static` 静态文件
- `templates` 模板文件

## .env

先将`.env.example`改名为`.env`
```text
当plugin为0的时候，DuckyGo会关闭MySQL和Redis的连接，就可以自己单独运行 不会报错.
```

## 使用说明
- `/static`就是`static`文件夹里面的内容
