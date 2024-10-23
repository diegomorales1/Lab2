FROM golang:1.23.0

WORKDIR /Lab2

COPY . .

RUN go mod tidy
RUN go build -o NodoTai .

EXPOSE 50071

CMD ["./NodoTai"]
