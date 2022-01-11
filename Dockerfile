# 使用官方 ubuntu
FROM ubuntu

# 定义工作目录
WORKDIR /usr/src/app

# 将本地代码复制到工作目录内
COPY . ./

# 启动服务
CMD ["./run"]

EXPOSE 8080
