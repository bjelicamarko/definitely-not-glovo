FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ./ReviewService

EXPOSE 8085

CMD [ "./ReviewService" ] 
