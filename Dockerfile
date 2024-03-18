FROM golang:1.21.8-alpine

LABEL maintainer="Nizigama"

WORKDIR /opt/auction-api

COPY . .

EXPOSE 3000

ENTRYPOINT ["go", "run", "./..."]
