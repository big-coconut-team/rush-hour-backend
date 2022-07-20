FROM golang:1.18-alpine AS build

ADD ./src /app
WORKDIR /app

RUN apk update
RUN apk add build-base

RUN go mod download

RUN go build -tags musl -o /product_svc

FROM alpine:3.16.0

COPY --from=build /product_svc /product_svc
COPY --from=build /app/.env /.env

EXPOSE 8001
CMD [ "/product_svc" ]
