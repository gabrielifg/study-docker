FROM debian

RUN apt-get update && apt-get install -y apache2
RUN chown www-data:www-data /var/lock && chown www-data:www-data /var/run/ && chown www-data:www-data /var/log/
ENV APCHE_LOCK_DIR="/var/lock"
ENV APCHE_PID_FILE="/var/run/apache2.pid"
ENV APCHE_RUN_USER="www-data"
ENV APCHE_RUN_GROUP="www-data"
ENV APCHE_LOG_DIR="/var/log/apache2"

# COPY index.html /var/www/html/
ADD index.html /var/www/html

LABEL description="Webserver"
LABEL version="1.0.0"

USER root

WORKDIR /var/www/html/

VOLUME /var/www/html/
EXPOSE 80

ENTRYPOINT [ "usr/sbin/apachectl" ]
CMD [ "-D", "FOREGROUND" ]
