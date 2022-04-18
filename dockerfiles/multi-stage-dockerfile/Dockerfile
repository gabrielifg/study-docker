FROM golang AS build

WORKDIR /app

ADD . /app
RUN go mod init github.com/gabrielifg/meugo
RUN go build -o meugo

FROM alpine

WORKDIR /meugo
COPY --from=build /app/meugo /meugo/
ENTRYPOINT ./meugo