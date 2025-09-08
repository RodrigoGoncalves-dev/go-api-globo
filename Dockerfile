FROM golang:latest as build

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd

FROM alpine:latest  

WORKDIR /app

RUN apk add --no-cache bash curl

COPY --from=build /app/api ./

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

EXPOSE 5000

ENTRYPOINT ["/wait-for-it.sh", "postgres:5432", "--", "./api"]
