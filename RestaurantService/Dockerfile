FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ./RestaurantService

EXPOSE 8082

CMD [ "./RestaurantService" ] 