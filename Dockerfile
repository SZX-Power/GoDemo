FROM golang:1.16.12-alpine
ENV GO111MODULE=on \
    GOPROXY = https://goproxy.cn,direct \
ENV TZ Asia/Shanghai
WORKDIR /app
COPY . .
RUN go build ./main.go
ENTRYPOINT ["/app/main"]
