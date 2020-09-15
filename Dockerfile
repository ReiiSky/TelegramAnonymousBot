FROM golang:1.15.0-alpine3.12

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go get
EXPOSE 6007
CMD ["make", "run"]