FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app 
COPY --from=build /app/main .
ENV PORT=${PORT}
ENV QDRANT_ADDR=${QDRANT_ADDR}
CMD [ "./main"]