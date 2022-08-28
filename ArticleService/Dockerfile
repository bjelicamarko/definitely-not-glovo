FROM golang:1.18-alpine as build

WORKDIR /app

COPY . .

RUN go build -o ./ArticleService

EXPOSE 8083

CMD [ "./ArticleService" ] 