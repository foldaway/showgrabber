FROM golang:1-alpine

WORKDIR /app

RUN apk add --update nodejs yarn

CMD ["./docker-cmd.sh"]
