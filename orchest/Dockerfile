FROM golang:1.18-alpine AS build

ADD ./src /app
WORKDIR /app

RUN apk update
RUN apk add build-base

RUN go mod download

RUN go build -tags musl -o /orchest_svc

FROM alpine:3.16.0

COPY --from=build /orchest_svc /orchest_svc

EXPOSE 3333
CMD [ "/orchest_svc" ]
