FROM registry.cn-hangzhou.aliyuncs.com/twilikiss/golang:1.22 AS build-stage
WORKDIR /app
COPY . ./
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build  -o /shifu-demo shifu-demo/main.go
FROM registry.cn-hangzhou.aliyuncs.com/twilikiss/distroless_base-debian:1.0 AS build-release-stage
WORKDIR /
COPY --from=build-stage /shifu-demo /shifu-demo
COPY --from=build-stage /app/shifu-demo/etc/config.toml /etc/config.toml
USER nonroot:nonroot
ENTRYPOINT ["/shifu-demo"]