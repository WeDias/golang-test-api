FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN apk update && apk add bash && bash setup-env.sh docker
RUN go build -o run-app

EXPOSE 3000

CMD [ "./run-app" ]