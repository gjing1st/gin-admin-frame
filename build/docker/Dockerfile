# 多阶段构建
#构建一个 builder 镜像，目的是在其中编译出可执行文件mck
#构建时需要将此文件放到代码根目录下
FROM golang:alpine  as builder
#将上层整个文件夹拷贝到/build
ADD . /build/src
WORKDIR /build/src
#去掉了调试信息 -ldflags="-s -w" 以减小镜像尺寸
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go mod tidy \
    && go build -ldflags="-s -w"  -o gaf /build/src/cmd/gaf/main.go

FROM alpine
#更新软件源  使用挂载卷挂载 /etc/localtime:/etc/localtime做宿主机时间映射
RUN #apk update --no-cache && apk add --no-cache tzdata \
# && apk add --no-cache docker-cli \
# && apk add --no-cache  docker-compose
##设置本地时区，这样我们在日志里看到的是北京时间了
#ENV TZ Asia/Shanghai


COPY --from=builder  /build/src/config/config.yml /home/config/config.yml
COPY --from=builder /build/src/gaf /home/gaf
WORKDIR /home
CMD ["./gaf"]
EXPOSE 9680

#docker run  -d --name publish -p 9680:9680 -v /etc/localtime:/etc/localtime publish
# docker run -d --pid=host --privileged=true -p 9680:9680 -v /etc/sysconfig/network-scripts:/etc/sysconfig/network-scripts -v /etc/localtime:/etc/localtime github.com/gjing1st/gin-admin-frame
