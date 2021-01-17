# Travel

数据库大作业：旅游管理系统

后端：Golang (gin + gorm)

前端：Vue (ant design of vue)



### 功能

+ [x] 航班，出租车，宾馆房间和客户基础数据的入库，更新。
+ [x] 预定航班，出租车，宾馆房间。
+ [x] 查询航班，出租车，宾馆房间，客户和预订信息。

+ [x] 查询某个客户的旅行线路。



### 注意：

+ 文件夹中有travel.exe（windows上的可执行文件）

+ 默认8080端口
+ 需运行init.sql创建travel数据库

+ 可用cmd使用-init参数初始化

  + ```shel
    travel.exe -init
    ```

+ 初始化的用户名与密码为admin，123456
+ 只有名为admin才有管理员权限
+ 若不使用-init参数初始化则需要自己注册管理员和用户



### 不使用可执行文件启动：

### 1.获取代码

```shell
git clone https://github.com/SSunSShine/travel
```

### 2.下载依赖

```shell
go mod tidy
```

### 3.修改配置信息

```shell
vim ./conf/configuration.yaml
```

### 4.初始化并运行

```shell
init.sql // 创建travel数据库

go run ./ -init
```

