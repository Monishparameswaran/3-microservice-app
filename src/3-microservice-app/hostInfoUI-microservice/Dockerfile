FROM golang:alpine as builder
WORKDIR /app
COPY /static .
COPY go.mod .
COPY main.go .
RUN go mod download && \
    go build main.go
EXPOSE 8084
ENTRYPOINT [ "./main" ]