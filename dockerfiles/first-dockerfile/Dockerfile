FROM debian

RUN apt-get update && apt-get install -y apache2
ENV APCHE_LOCK_DIR="/var/lock"
ENV APCHE_PID_FILE="/var/run/apache2.pid"
ENV APCHE_RUN_USER="www-data"
ENV APCHE_RUN_GROUP="www-data"
ENV APCHE_LOG_DIR="/var/log/apache2"

LABEL description="Webserver"
LABEL version="1.0.0"

VOLUME /var/www/html/
EXPOSE 80

## instrução a baixo para de fato executar o apache e não somente cria-ló
## Modo exec (recomendado pela documentação)

ENTRYPOINT [ "usr/sbin/apachectl" ]
CMD [ "-D", "FOREGROUND" ]

## Modo shell (não pode ser usado junto com o modo exec)
## CMD /usr/sbin/apachectl -D FOREGROUND

## OBS: Instrução utilizada para dizer ao apache que ele seja executado em 1° plano e não em modo daemon