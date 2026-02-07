# 介绍

简单的图床，使用golang+vue开发，采用BadgerDB作为存储图片元数据，图片存储在本地目录

功能局限性：

- 只支持单用户
- 只支持图片本地存储
- 只支持上传jpg/jpeg/png/gif/webp，通过图片解码避免图片伪造
- jpg/jpeg/png自动转码为webp格式，根据图片分辨率自动设置webp压缩质量
- 上传图片大小硬编码限制100M

**项目只适合个人使用，不考虑对接云存储，不考虑支持多用户**

# 部署运行

```shell
./fast-image 
```

首次运行会自动在可执行文件所在目录创建data/config.yaml

```yaml
address: 0.0.0.0:8080
account: admin
password: admin
```

可以通过设置环境变量，覆盖配置

```shell
FI_ADDRESS=0.0.0.0:8080
FI_ACCOUNT=admin
FI_PASSWORD=admin
```

# 程序数据目录说明

```shell
# 程序配置文件，会自动创建
data/config.yaml

# BadgerDB持久化存储，此目录存储图片的元数据信息
data/badger/

# HTTP请求日志，忽略了上传图片和请求图片的日志
data/http_logs/

# 上传图片存储的目录，会按照YYYY-MM-DD目录存储图片
data/images/YYYY-MM-DD

# 程序运行日志
data/logs/
```