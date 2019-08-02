FROM golang:1.12 AS builder

ENV GOPROXY https://goproxy.io
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o gosharp
#RUN GOOS=linux GOARCH=amd64 go build -v -o /app/xchef-dashboard/sso-server


FROM alpine

WORKDIR /app

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/gosharp .
COPY --from=builder /app/config/conf.yaml ./config/conf.yaml
COPY --from=builder /app/docs/swagger.json ./docs/gosharp.json

#EXPOSE 8080
CMD [ "/app/gosharp" ]