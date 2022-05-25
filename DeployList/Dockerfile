FROM golang:1.17.0-alpine
ENV GO111MODULE=on \
    GOPROXY=https://mirrors.aliyun.com/goproxy/
ENV TZ Asia/Shanghai
WORKDIR /app
COPY . .
RUN go build ./main.go
ENTRYPOINT ["/app/main"]
