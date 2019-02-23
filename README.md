<a href="https://coding13.com/"><img height="160" src="/frankxi/league/blob/master/comm/images/logo.png?raw=true"></a>

Official site: [http://www.coding13.com](http://www.coding13.com/)
feature
-- 自动生成业务代码
-- 支持跨域
-- 支持JWT权限访问
-- 支持接口验签
-- ORM 采用 Gorm
-- 支持redis集成

目录结构
-- biz    //业务
-- comm   //公共
-- docs   //文档
-- gen    //生成器
-- microservices // 微服务
-- middleware // 中间件
-- routers  // http 路由配置
-- utils //工具类
-- main.go  //启动入口



依赖
github.com/go-ini/ini

go get github.com/gin-gonic/gin github.com/gin-contrib/cors
github.com/gin-gonic/gin
github.com/gin-contrib/cors

github.com/swaggo/gin-swagger
github.com/swaggo/gin-swagger/swaggerFiles

github.com/dgrijalva/jwt-go

github.com/go-sql-driver/mysql

github.com/jinzhu/gorm
github.com/gomodule/redigo/redis

-- log
github.com/sirupsen/logrus
