# Stage 1 (to create a "build" image, ~850MB)
FROM golang:1-alpine3.17 AS builder

COPY . /api/
WORKDIR /api/
RUN set -Eeux && \
    go mod download && \
    go mod verify

RUN GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-w -s -X 'main.Version=1.0'" \
    -o app main.go
RUN go test -cover -v ./...

# Stage 2 (to create a downsized "container executable", ~5MB)

# If you need SSL certificates for HTTPS, replace `FROM SCRATCH` with:
#
#   FROM alpine:3.17.1
#   RUN apk --no-cache add ca-certificates
#
FROM scratch
WORKDIR /root/
COPY --from=builder /api .

EXPOSE 8080
ENTRYPOINT ["./app"]