FROM golang:1.18-alpine AS build

ADD ./src /app
WORKDIR /app

RUN apk update
RUN apk add build-base

RUN go mod download

RUN go build -o /backend_con

FROM alpine:3.16.0

COPY --from=build /backend /backend_con
COPY --from=build /app/.env /.env

EXPOSE 8088
CMD [ "/backend_con" ]
