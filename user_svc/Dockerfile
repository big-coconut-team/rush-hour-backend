FROM golang:1.18-alpine

ADD ./src /app
WORKDIR /app

RUN go mod download

RUN go build -o /user_svc

EXPOSE 8000
CMD [ "/user_svc" ]
