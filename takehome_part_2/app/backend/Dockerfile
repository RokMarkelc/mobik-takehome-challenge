FROM golang:1.25

WORKDIR /app

# Copy source
COPY . .

# Preload modules
RUN go mod download && go mod tidy

# Build the app
RUN go build -o ./backend.go ./cmd/server

EXPOSE 8080

CMD ["./backend.go"]
