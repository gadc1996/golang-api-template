FROM golang:1.19-alpine

WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

EXPOSE 8080

# CMD ["air", "-c", ".air.toml"]
CMD ["go", "run", "main.go"]