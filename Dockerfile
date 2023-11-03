FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN go build -o /go-rest-api

ARG APP_PORT
ENV APP_PORT=$APP_PORT

EXPOSE $APP_PORT

CMD [ "/go-rest-api" ]