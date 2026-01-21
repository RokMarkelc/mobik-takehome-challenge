FROM golang:1.22

WORKDIR /app

# Preload modules
COPY go.mod ./
RUN go mod download

# Install air for hot reload
RUN go install github.com/cosmtrek/air@latest

# Copy source
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
