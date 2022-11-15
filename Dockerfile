# Build stage
FROM golang:1.18-alpine as builder
RUN mkdir -p /app
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -buildvcs=false -o=appbin ./cmd/server/

# Deploy stage
FROM alpine
RUN mkdir -p /app
WORKDIR /app
COPY --chown=0:0 --from=builder /app/ ./
ENTRYPOINT ["/app/appbin"]