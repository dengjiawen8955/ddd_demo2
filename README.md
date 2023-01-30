# dc2

## 快速开始

启动 env MySQL 和 redis 依赖

```bash
make up.env
```

初始化数据库 (需要修改 config.yaml 中的数据库配置)

```bash
make exec.sql
```

编译

```bash
make gen
make tidy
make build
```

命令行启动

```bash
make start
```

docker 后台启动

```bash
make up
```

## 问题和解决

### docker 网络问题

docker 里面无法访问宿主机的网络, 所以需要修改 docker-compose.yml 中的网络配置

## 工具

stogo

yaml2go

go2json

https://www.go2json.com/

json2go

