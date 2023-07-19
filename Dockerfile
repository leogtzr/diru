FROM golang:latest

RUN apt-get update && apt-get install -y redis-server

WORKDIR /app

COPY . .

RUN go build -o surl

CMD ["bash", "-c", "redis-server & ./surl"]

