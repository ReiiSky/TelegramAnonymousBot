FROM golang:1.15.0-alpine3.12

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get
CMD ["make", "run"]