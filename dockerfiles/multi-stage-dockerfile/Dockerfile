FROM golang

WORKDIR /app

ADD . /app
RUN go mod init github.com/gabrielifg/meugo
RUN go build -o meugo

ENTRYPOINT ./meugo