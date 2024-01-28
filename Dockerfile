FROM golang:1.19.0-alpine
WORKDIR /app/src
ADD go.mod ../
ADD src .
RUN go build -o ../app .
CMD ["../app"]