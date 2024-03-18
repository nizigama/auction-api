FROM golang:1.21.8-alpine

LABEL maintainer="Nizigama"

WORKDIR /opt/auction-api

COPY . .

RUN go build && chmod +x ./Web3AuctionApi

EXPOSE 3000

ENTRYPOINT ["./Web3AuctionApi"]
