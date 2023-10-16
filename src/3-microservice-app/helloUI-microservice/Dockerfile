FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download && \
    go build main.go

FROM istio/distroless:1.19-2023-10-13T03-27-30
WORKDIR /myapp
COPY --from=builder /app/main ./
EXPOSE 8082
ENTRYPOINT [ "./main" ]