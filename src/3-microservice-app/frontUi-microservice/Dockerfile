FROM golang:alpine as builder
WORKDIR /app
COPY . .
RUN go mod download 
RUN go build main.go
EXPOSE 8080
CMD [ "./main" ]