FROM golang:1.19.0-alpine
WORKDIR /app
ADD go.mod ./
ADD src ./src
ADD public ./public
RUN go build -o app ./src
CMD ["./app"]