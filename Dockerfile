FROM golang:latest

RUN apt-get update && apt-get install -y redis-server

WORKDIR /app

COPY . .

RUN go build -o surl

# Add entry point script
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Use entry point script
CMD ["/app/entrypoint.sh"]
