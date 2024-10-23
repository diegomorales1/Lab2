FROM golang:1.23.0

WORKDIR /Lab2

COPY . .

RUN go mod tidy
RUN go build -o datanode .

EXPOSE 50054

CMD ["./datanode"]
