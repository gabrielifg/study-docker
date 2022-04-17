FROM debian

RUN apt-get update && apt-get install -y apache2 && apt-get clean
ENV APCHE_LOCK_DIR="/var/lock"
ENV APCHE_PID_FILE="/var/run/apache2.pid"
ENV APCHE_RUN_USER="www-data"
ENV APCHE_RUN_GROUP="www-data"
ENV APCHE_LOG_DIR="/var/log/apache2"

LABEL description="Webserver"
LABEL version="1.0.0"

VOLUME /var/www/html/
EXPOSE 80