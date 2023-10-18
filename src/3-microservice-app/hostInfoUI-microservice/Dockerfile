FROM golang:alpine as builder
WORKDIR /app
COPY /static .
COPY go.mod .
COPY main.go .
RUN go mod download && \
    go build main.go
FROM alpine:3.18.4
COPY --from=builder /app/main .
COPY /static .
EXPOSE 8084
ENTRYPOINT [ "./main" ]