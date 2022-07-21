FROM golang:1.18-alpine AS build

ADD ./src /app
WORKDIR /app

RUN apk update
RUN apk add build-base

RUN go mod download

RUN go build -tags musl -o /backend

FROM alpine:3.16.0

COPY --from=build /backend /backend

EXPOSE 8088
CMD [ "/backend" ]
