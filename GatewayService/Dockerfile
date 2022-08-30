FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ./GatewayService

EXPOSE 8080

CMD [ "./GatewayService" ] 