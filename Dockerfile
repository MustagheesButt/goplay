FROM golang

EXPOSE 8080

WORKDIR /app

COPY . .

RUN go build

CMD ["./goplay"]