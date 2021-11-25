FROM golang:alpine as builder


RUN apk update && apk add --no-cache git 

WORKDIR /app


COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates && apk add --no-cache tzdata && apk --no-cache add curl

WORKDIR /root/

COPY --from=builder /app/main .

# Environment
ENV MYSQL_USERNAME 'dev'
ENV MYSQL_PASSWORD "secret"
ENV MYSQL_DATABASE "air-asia"
ENV MYSQL_HOST "docker.for.mac.host.internal"
ENV MYSQL_PORT "3406"

ENV SFTP_USER "RM8005-jeffrey"
ENV SFTP_PASS "4080"
ENV SFTP_HOST "192.168.0.198"

ENV MAILJET_KEY "42b6094a466608288b8ba50b32509da6"
ENV MAILJET_SECRET "bccdaa44b4a44acb59b1c5e36e85aaa4"

ENV EMAIL "leanwf1117@gmail.com"
ENV REDISHOST "localhost:6379"

CMD ["./main"]