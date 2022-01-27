FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/fake-data-service .
RUN rm -rf /app/internal
RUN rm -rf /app/pkg
RUN rm -rf /app/test
RUN rm -rf /app/main.go
RUN rm -rf /app/go.mod
EXPOSE 8000

CMD [ "/app/fake-data-service" ]
