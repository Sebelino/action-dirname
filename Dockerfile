FROM golang:1.19.1
WORKDIR /build/
COPY app.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.16.2
WORKDIR /app/
COPY --from=0 /build/app ./
ENTRYPOINT ["./app"]
