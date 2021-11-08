FROM ubuntu:20.04

ARG IROHAD_BIN
ARG IROHA_CLI_BIN

COPY ${IROHAD_BIN} /usr/bin/irohad
COPY ${IROHA_CLI_BIN} /usr/bin/iroha-cli

WORKDIR /opt/iroha_data

COPY entrypoint.sh wait-for-it.sh /
RUN chmod +x /entrypoint.sh /wait-for-it.sh
ENTRYPOINT ["/entrypoint.sh"]
CMD ["irohad"]
