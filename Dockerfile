FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -v -o server

FROM debian:buster-slim
COPY --from=builder /app/server /server
CMD ["/server"]