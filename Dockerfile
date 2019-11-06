FROM golang:latest AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.io
RUN go mod vendor && \
    go build cmd/main.go

FROM debian:latest
RUN mkdir /config
COPY --from=build /app/main .
COPY --from=build /app/config/config.yml /config/
ENTRYPOINT ["./main"]
EXPOSE 8888
