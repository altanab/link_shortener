FROM golang:latest AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN go build -o main main.go

EXPOSE 9080
CMD /app/maingit