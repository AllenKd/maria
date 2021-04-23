FROM golang as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o stress_test ./cmd/container

FROM ubuntu

WORKDIR /app
COPY --from=builder /app/stress_test ./
COPY --from=builder /etc/ssl /etc/ssl

COPY ./configs/config.yaml ./configs/

RUN apt-get update \
    && apt-get install -y \
        nmap \
        vim

# Command to run the executable
CMD ["sleep", "infinity"]
