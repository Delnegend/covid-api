FROM golang:alpine AS builder

WORKDIR /build
COPY ./utils/* ./utils/
COPY main.go go.mod go.sum .
RUN go build -o main .

FROM alpine:latest

COPY --from=builder /build/main /build/main
EXPOSE 8080

CMD ["/build/main"]