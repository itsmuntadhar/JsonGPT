FROM golang:alpine as builder

WORKDIR /src
RUN apk add --no-cache ca-certificates
COPY ../src /src
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/jsongpt
FROM alpine:3.15

COPY --from=builder /bin/jsongpt /app/jsongpt
RUN apk add --no-cache ca-certificates bash
WORKDIR /app
