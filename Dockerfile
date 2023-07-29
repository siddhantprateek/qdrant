FROM golang:1.20-alpine AS build 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN go build -tags netgo -ldflags '-s -w' -o main

FROM alpine:latest
WORKDIR /app 
COPY --from=build /app/main .
ENV PORT=${PORT}
ENV QDRANT_ADDR=${QDRANT_ADDR}
CMD [ "./main"]