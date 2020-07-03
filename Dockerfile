FROM golang:1.14

WORKDIR /app
COPY . .
RUN go build -o reactor .

EXPOSE 8080

CMD ["/app/reactor"]
