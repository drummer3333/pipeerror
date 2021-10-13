FROM golang:buster

WORKDIR /app

ADD . .
RUN chmod +x docker-entrypoint.sh

ENTRYPOINT [ "./docker-entrypoint.sh" ]
