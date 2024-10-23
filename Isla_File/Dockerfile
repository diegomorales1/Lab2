FROM golang:1.23.0

WORKDIR /Lab2

COPY . .

RUN go mod tidy
RUN go build -o Isla_File .

EXPOSE 50051

CMD ["./Isla_File"]
