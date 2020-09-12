FROM golang:latest

WORKDIR /app

COPY . .

ENV GOPROXY="https://goproxy.cn"

RUN go mod download

RUN go build -o meido .

EXPOSE 8080

CMD [ "./meido" ]
