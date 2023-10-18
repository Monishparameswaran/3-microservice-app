FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download && \
    go build main.go
EXPOSE 8082
ENTRYPOINT [ "./main" ]