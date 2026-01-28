FROM golang:1.25

WORKDIR /app

# Copy source
COPY . .

# Preload modules
RUN go mod download && go mod tidy

# Build the app
RUN go build -o /app/backend.go /app/cmd/server

EXPOSE 8080

ENTRYPOINT ["/app/docker/entrypoint.sh"]
CMD ["/app/backend.go"]
