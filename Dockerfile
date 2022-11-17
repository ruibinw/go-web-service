# Build stage
FROM golang:1.19.3-alpine as builder
RUN apk add build-base #to provide gcc for building sqlite
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -buildvcs=false -o=appbin ./

# Deploy stage
FROM alpine
RUN mkdir -p /app
WORKDIR /app
COPY --chown=0:0 --from=builder /app/ ./
ENTRYPOINT ["/app/appbin"]