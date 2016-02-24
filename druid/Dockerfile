FROM java:8

ENV DRUID_VERSION 0.8.3
RUN wget -q -O - http://static.druid.io/artifacts/releases/druid-$DRUID_VERSION-bin.tar.gz | tar -xzf - -C /opt

RUN mv /opt/druid-$DRUID_VERSION /opt/druid

VOLUME ["/data/druid"]

WORKDIR /opt/druid

COPY config /opt/druid/config
COPY bin /opt/druid/bin
COPY specs /opt/druid/specs
COPY entrypoint.sh /opt/druid/entrypoint.sh

# pull maven dependencies
RUN java -classpath "config/_common:lib/*" io.druid.cli.Main tools pull-deps

ENTRYPOINT ["./entrypoint.sh"]
