# build the backend using Builder container
FROM golang:1.23.5 AS BUILDER

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o backend .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/backend /app/backend

# Set the working directory
WORKDIR /app

# Expose the port that the backend app runs on
EXPOSE 8787

# Run the backend application
CMD ["./backend"]
