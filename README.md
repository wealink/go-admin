go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod init

# 进入项目目录
cd go-admin

# 安装依赖
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod init

# mysql数据结构
./docs/mysql.sql

# 配置文件
./config/settings.yml

# 启动服务
go run main.go
```