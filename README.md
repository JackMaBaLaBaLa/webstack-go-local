## 源自：https://github.com/ch3nnn/webstack-go

**仅针对docker部署方式，直接部署未测试。**

将css和js资源放在本地，可不依赖于网络，从而可以在内网机器都不能连外网的情况下部署和使用。

### 1、创建镜像(在能连外网的机器上创建镜像)

a.编辑配置文件

编辑configs/docker_configs.toml，改为要部署的环境配置信息

然后执行：

```shell
cd webstack-go-local-main

docker build -t webstack-go:v1.0.0 .
```

b.导出镜像

```shell
docker save -o webstack-go:v1.0.0.tar.gz webstack-go:v1.0.0
```

c.导入镜像

将webstack-go:v1.0.0.tar.gz导入待部署的内部环境

```shell
docker load -i webstack-go:v1.0.0.tar.gz
```

d.依次再导入redis和mysql镜像

redis和mysql镜像可以在能连外网的机器上执行以下命令

```shell
docker pull mysql/mysql-server:5.7

docker pull redis:6.2.4
```

按照以上方法导入内网环境即可

### 2.运行

修改docker-compose.yml-------->若有需要

```shell
cd webstack-go-local-main

docker-compose up -d #或者 docker-compose -f docker-compose.yml up -d
```

### 3.访问

http://ip:9999/install 进行项目初始化

初始化完成后，可正常使用

http://ip:9999
