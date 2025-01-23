# build the backend using Builder container
FROM golang:1.23.5

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o backend ./cmd/server/main.go

# TODO: remove source code after build optional

# Expose the port that the backend app runs on
EXPOSE 8787

# Run the backend application
CMD ["./backend"]
