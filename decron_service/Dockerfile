FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
