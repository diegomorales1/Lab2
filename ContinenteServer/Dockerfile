FROM golang:1.23.0

WORKDIR /Lab2

COPY . .

RUN go mod tidy
RUN go build -o ContinenteServer .

EXPOSE 50053

CMD ["./ContinenteServer"]
