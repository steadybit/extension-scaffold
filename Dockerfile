# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19-alpine AS build

WORKDIR /app

RUN apk add build-base
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN make build

##
## Runtime
##
FROM alpine:3.16

WORKDIR /

COPY --from=build /app/extension /extension

EXPOSE 8080

ENTRYPOINT ["/extension"]
