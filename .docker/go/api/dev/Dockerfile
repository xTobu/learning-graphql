# FROM golang:1.11.2
# WORKDIR /root/go
# COPY . .
# RUN go build -o ApiServer -ldflags "-X main.ENV=development"


# FROM debian:9-slim as stage
# WORKDIR /root/
# COPY --from=builder ./root/go/ApiServer ./
# COPY --from=builder ./root/go/config ./config/

# EXPOSE 3000

# CMD ["/root/ApiServer"]

FROM golang:1.11.2
WORKDIR /root/go
RUN pwd && ls -al
COPY . .
RUN go get github.com/oxequa/realize
CMD ["realize", "start"]
EXPOSE 3000

# RUN 执行命令并创建新的镜像层，RUN 经常用于安装软件包。

# CMD 设置容器启动后默认执行的命令及其参数，但 CMD 能够被 docker run 后面跟的命令行参数替换。

# ENTRYPOINT 配置容器启动时运行的命令。