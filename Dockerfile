FROM golang:alpine AS builder
ADD . /go/src/gosharp
WORKDIR /go/src/gosharp
# install git
RUN apk add --no-cache git bzr
RUN go get golang.org/x/sys/unix
RUN go get -u -v github.com/kardianos/govendor && \
   govendor sync
RUN GOOS=linux GOARCH=amd64 go build -v -o /go/src/gosharp/robo-server

FROM alpine
WORKDIR /root
RUN apk add -U tzdata && \
   ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
COPY --from=builder /go/src/gosharp .
#EXPOSE 8080
CMD [ "./robo-server" ]