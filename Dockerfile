FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -v -o storiapp

FROM golang:1.19
WORKDIR /app
COPY --from=builder /app/storiapp /app/storiapp
COPY test1.csv /app/test1.csv
CMD ["./storiapp"]