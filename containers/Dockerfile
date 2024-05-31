FROM golang:1.22 AS builder

RUN mkdir /distributing
WORKDIR /distributing
COPY notify/ notify/
COPY server/ server/
WORKDIR /distributing/server
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -tags=containers

FROM scratch
WORKDIR /
COPY --from=builder /distributing/server .
COPY ./config.toml .
EXPOSE 8080/tcp
CMD ["/server"]