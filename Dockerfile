# 使用较小的 Debian 作为基础镜像
# 外网下镜像换成abbhb/pdftopng_system:latest即可
FROM 10.15.247.193/test/pdftopng_system:latest

# 将工作目录设置为 /app
WORKDIR /app
ENV GOPROXY https://goproxy.cn

ENV GO111MODULE on
ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0

# 复制当前目录内容到容器中的 /app 目录
COPY . /app

RUN go build -v -ldflags="-w -s" -o pdf2png-node .


# 指定启动命令
CMD ["./pdf2png-node"]
