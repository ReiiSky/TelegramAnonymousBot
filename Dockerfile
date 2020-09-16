FROM golang:1.15.0-alpine3.12

RUN mkdir /app
COPY . /app
WORKDIR /app
EXPOSE 6007
CMD ["go", "run", "cmd/server/main.go"]
