FROM golang:alpine as builder

RUN apk --no-cache add git

WORKDIR /data0/apps

RUN git clone https://github.com/snail2sky/http-echo.git 

WORKDIR /data0/apps/http-echo

RUN go build

FROM alpine:latest

WORKDIR /data0/apps/http-echo

COPY --from=builder /data0/apps/http-echo/http-echo .

ENTRYPOINT [ "./http-echo" ]

