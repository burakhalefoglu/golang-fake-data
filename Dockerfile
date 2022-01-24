FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/fake-person-worker-service .
RUN rm -rf /app/createFakePerson
RUN rm -rf /app/database
RUN rm -rf /app/fakePersonStruct
RUN rm -rf /app/main.go
RUN rm -rf /app/helper
RUN rm -rf /app/mongodb
RUN rm -rf /app/Dockerfile
RUN rm -rf /app/go.mod


CMD [ "/app/fake-person-worker-service" ]
