FROM golang:1.23-alpine

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go mod tidy
