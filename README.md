### 目录结构
* app 路由
* config 配置文件
* handler 控制中心
* db 连接数据库
* models 操作数据库
* response 返回数据结构
* utils 工具类
* config.yaml 配置文件
* main.go 入口文件

### 保存表情包报错
windows下的为my.ini(linux下的为my.cnf)
```
[client]
default-character-set = utf8mb4
[mysql]
default-character-set = utf8mb4
[mysqld]
character-set-client-handshake = FALSE
character-set-server = utf8mb4
collation-server = utf8mb4_unicode_ci
init_connect='SET NAMES utf8mb4'

alter table TABLE_NAME convert to character set utf8mb4 collate utf8mb4_bin;
```



mac打包linux
```
export CGO_ENABLED=1
export GOOS=linux
export GOARCH=amd64
export CC=/usr/local/gcc-4.8.1-for-linux64/bin/x86_64-pc-linux-gcc
go build -o ./bin/example ./cmd/
```

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

windows打包linux
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```

发布到服务器
```
$ ssh root@47.104.198.49
```

修改代码权限
```
chmod -R 777 main
```

一直在后台运行
配置nohup文档
> nohup ./golangFile &

查看进程
> ps -aux|grep main(文件名)

杀死进程
> kill -9 main

### WIN10快捷键
```
创建新的虚拟桌面：Win + Ctrl + D
关闭当前虚拟桌面：Win + Ctrl + F4
切换虚拟桌面：Win + Ctrl +左/右
```

依赖包安装路径
/Users/admin/go/src/github.com

[go编译报错参考文档](https://xudeyi.com/2019/02/25/mac-build-go-sqlite3-error/)

错误信息
```
sqlite3-binding.c:32753:42: error: use of undeclared identifier 'pread64'
sqlite3-binding.c:32771:42: error: use of undeclared identifier 'pwrite64'
sqlite3-binding.c:32903:22: error: invalid application of 'sizeof' to an incomplete type 'struct unix_syscall []'
sqlite3-binding.c:32912:22: error: invalid application of 'sizeof' to an incomplete type 'struct unix_syscall []'
sqlite3-binding.c:32939:20: error: invalid application of 'sizeof' to an incomplete type 'struct unix_syscall []'
sqlite3-binding.c:32956:16: error: invalid application of 'sizeof' to an incomplete type 'struct unix_syscall []'
sqlite3-binding.c:14066:38: note: expanded from macro 'ArraySize'
sqlite3-binding.c:32960:14: error: invalid application of 'sizeof' to an incomplete type 'struct unix_syscall []'
sqlite3-binding.c:14066:38: note: expanded from macro 'ArraySize'
sqlite3-binding.c:35613:11: warning: type specifier missing, defaults to 'int' [-Wimplicit-int]
sqlite3-binding.c:32757:49: note: expanded from macro 'osPread64'
sqlite3-binding.c:35725:17: warning: type specifier missing, defaults to 'int' [-Wimplicit-int]
sqlite3-binding.c:32775:57: note: expanded from macro 'osPwrite64'
```

问题
Mac未安装linux的交叉编译器。解决方法，安装linux的交叉编译器

下载地址
Compiling for Linux 32 and Linux 64 on MacOS X

多表查询
```go
model := Db().Model(&Order{}).Select("orders.*").Order("orders.create_time desc").Joins("left join users on orders.user_id = users.id").Where("orders.id LIKE ? AND orders.create_time BETWEEN ? AND ? AND orders.status IN (?)", "%"+orderID+"%", start, end, status)
if len(links) > 0 {
    model = model.Where("users.link_id IN (?)", links)
}
model.Count(&total).Limit(pageSize).Offset((current - 1) * pageSize).Preload("Commodity").Preload("Address").Find(&order)
row := model.Select("sum(orders.total_price)").Row()
row.Scan(&count)
row1 := model.Where("orders.status IN (?)", []string{"2", "4", "5", "6"}).Select("sum(orders.total_price)").Row()
row1.Scan(&inc)
```

