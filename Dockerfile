FROM golang:1.19.1 AS builder
WORKDIR /build/
COPY app.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.16.2
COPY --from=builder /build/app /app
ENTRYPOINT ["/app"]
