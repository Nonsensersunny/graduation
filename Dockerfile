FROM golang:latest AS build
RUN mkdir -p /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on
RUN go mod vendor && \
    go build cmd/main.go

FROM debian:latest
RUN mkdir /config
COPY --from=build /go/src/app/main .
COPY --from=build /go/src/app/config/config.yml /config/
ENTRYPOINT ["./main"]
EXPOSE 8888
