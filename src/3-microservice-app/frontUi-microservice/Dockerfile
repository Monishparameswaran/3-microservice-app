FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download 
RUN go build main.go
FROM alpine:3.18.4
WORKDIR /myapp
COPY --from=builder /app/main .
COPY index.html .
EXPOSE 8080
CMD [ "./main" ]