FROM golang:alpine as builder
WORKDIR /app
COPY ./static/ .
COPY main.go .
COPY go.mod .
RUN go mod download 
RUN go build main.go

FROM alpine:latest
WORKDIR /myapp
COPY --from=builder /app/main ./
EXPOSE 8080
ENTRYPOINT [ "./main" ]