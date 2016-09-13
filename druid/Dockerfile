FROM java:8

ENV DRUID_VERSION 0.9.1.1
RUN wget -q -O - http://static.druid.io/artifacts/releases/druid-$DRUID_VERSION-bin.tar.gz | tar -xzvf - -C /opt

RUN mv /opt/druid-$DRUID_VERSION /opt/druid

WORKDIR /opt/druid

# mysql metadata storage extension is not included in 0.9, so install separately
RUN mkdir -p extensions/mysql-metadata-storage

RUN wget -q -O - http://static.druid.io/artifacts/releases/mysql-metadata-storage-$DRUID_VERSION.tar.gz | \
  tar -xzvf - -C extensions

VOLUME ["/data/druid"]

COPY config config
COPY bin bin
COPY specs specs
COPY entrypoint.sh entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
