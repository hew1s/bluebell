# bluebell论坛项目

## 1010

### 使用搭建好的gin脚手架

### 用户表结构设计

### 基于雪花算法生成用户ID

### 基本思路

1.注册路由

2.controller控制器处理请求

​		1.获取前端传进来的参数

​		2.进行业务逻辑处理  由logic层进行业务逻辑处理

​				logic层进行业务逻辑的处理，涉及到数据库相关操作交由mysql层的程序进行增删改查操作

​		3.业务处理完返回响应

### 注册、登录、封装状态码

#### 注册

1.前端路由`signup`

2.控制器接收传来的数据`c.ShouldBindJSON(p)`

3.进行参数检测

4.交由业务逻辑层处理

​		1.检测用户是否已经存在（数据库查询username）

​		2.将新用户插入数据库

5.返回响应

#### 登录

1.前端路由`login`

2.控制器接收

3.进行参数检测

4.交由业务逻辑层进行处理

​		1.检测该用户是否存在（数据库查询username）

​				如果存在，将接收到的密码进行加盐与数据库中密码进行对比

5.返回响应

#### 封装状态码

将不同的状态码对应不同的msg信息

封装成一个程序

## 1011

### JWT（json web token）基于token的认证方式（JWT鉴权）

jwt包通过userid username等生成token信息，登录成功就将token信息传入前端，由前端保存在Authroization中

设计中间件，检验token是否存在请求头或者其他位置

#### access token  result token设计

暂定accesstoken 每十分钟刷新一次，只要result token还在就可以请求接口

（个人理解）防止有人获取到access token，只要将access token定时刷新就可以减少攻击

#### 某一时间只有一个设备登录

限制每个号在同一时间内只可在一台设备登录

个人解决方法（使用redis来保存 userID与token值对应的键值对）

每次用户登陆时，都会重新生成token，使用redis将userID与token相互对应以键值对存入redis数据库

当出现新token时，旧的token就会失效，而token校验程序会拿前端传来的token与现存redis中的token进行校验，

旧token不能通过校验所以强制退出

### 社区列表接口实现

路由注册

controller（如果有参数就接收参数进行校验）

logic实现业务逻辑

mysql进行数据库增删改查

controller返回响应

### 创建帖子功能实现、

首先创建对应结构体 一切以结构体操作

##### 帖子id，用户id，社区id，标题，内容，状态（可默认）

注册路由

controller

接收参数，通过token获取用户id 交由logic处理

logic实现具体业务，需要通过雪花算法来生成文章id

mysql将信息插入数据库

返回响应信息

### Makefile

借助`Makefile`我们在编译过程中不再需要每次手动输入编译的命令和编译的参数，可以极大简化项目编译过程。

## 问题

导入包循环引用