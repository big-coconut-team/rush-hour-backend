FROM golang:1.18-alpine AS build

ADD ./src /app
WORKDIR /app

RUN apk update
RUN apk add build-base

RUN go mod download

RUN go build -tags musl  -o /user_svc

FROM alpine:3.16.0

COPY --from=build /user_svc /user_svc
COPY --from=build /app/.env /.env

EXPOSE 8000
CMD [ "/user_svc" ]
