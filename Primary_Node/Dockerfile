FROM golang:1.23.0

WORKDIR /Lab2

COPY . .

RUN go mod tidy
RUN go build -o Primary_Node .

EXPOSE 50050 50060

CMD ["./Primary_Node"]
