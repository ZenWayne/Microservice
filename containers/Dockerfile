FROM golang:1.22 AS builder

RUN mkdir /distributing
WORKDIR /distributing
COPY server/ server/
WORKDIR /distributing/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -tags=containers

FROM scratch
WORKDIR /
COPY --from=builder /distributing/server .
COPY ./config.toml .
ENV MYSQL_URL=localhost
ENV MYSQL_PORT=3306
ENV MYSQL_USER=ethadmin
ENV MYSQL_PASSWORD=123
ENV MYSQL_DATABASE=emsvc
ENV REDIS_SERVERS=localhost:6379
ENV REDIS_PASS=123
EXPOSE 8080/tcp
CMD ["/server"]