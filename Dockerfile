FROM debian:9

ENV VPN_CLIENT_CONFIG_FILE=/tmp/myvpn-client-config

EXPOSE 8000

RUN apt-get -y update && apt-get -y install wget

COPY start.sh /start.sh
RUN chmod +x /start.sh

ENTRYPOINT /start.sh
