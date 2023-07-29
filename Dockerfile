FROM golang:1.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main
ENV PORT=${PORT}
ENV QDRANT_ADDR=${QDRANT_ADDR}
CMD ["./main"]