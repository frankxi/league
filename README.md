<a href="https://coding13.com/"><img height="160" src="https://raw.githubusercontent.com/frankxi/league/master/comm/images/logo.png"></a>

Official site: [http://www.coding13.com](http://www.coding13.com)

yyyy
1. feature
    * 自动生成业务代码
    * 支持跨域
    * 支持JWT权限访问
    * 支持接口验签
    * ORM 采用 Gorm
    * 支持redis集成
2. 目录结构
    * — biz //业务
    * — comm //公共
    * — docs //文档
    * — gen //生成器
    * — microservices // 微服务
    * — middleware // 中间件
    * — routers // http 路由配置
    * — utils //工具类
    * — main.go //启动入口

3. 业务代码生成
    * — gen //生成器
        * — 配置Config.json，配置数据库表、需要生成的table
        * — 运行gen.go 生成业务代码
        * — 配置routers 启动测试

4. 文档生成
    * — docs //文档生成
        * — install swaggo
        * — cd ./league
        * — swag init
        * — restart main.go
