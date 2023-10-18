FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download && \
    go build main.go

FROM alpine:3.18.4
WORKDIR /myapp
COPY --from=builder /app/main .
COPY /static .
EXPOSE 8082
ENTRYPOINT [ "./main" ]