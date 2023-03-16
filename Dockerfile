FROM golang:1.20-alpine

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o it-planet-task ./cmd/app/main.go

EXPOSE 8080

CMD [ "./it-planet-task" ]