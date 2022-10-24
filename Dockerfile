FROM golang:1.19.0
WORKDIR /home/wilton/repos/courses/go/trivia-tutorial

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy