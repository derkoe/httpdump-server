FROM golang:1.11-alpine
WORKDIR /go/src/github.com/derkoe/httpdump-server/
COPY httpdump-server.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o httpdump-server .

FROM scratch
COPY --from=0 /go/src/github.com/derkoe/httpdump-server/httpdump-server .
ENTRYPOINT ["/httpdump-server"]
EXPOSE 8080
