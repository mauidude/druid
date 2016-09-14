FROM java:8

ENV DRUID_VERSION 0.9.1.1
RUN wget -q -O - http://static.druid.io/artifacts/releases/druid-$DRUID_VERSION-bin.tar.gz | tar -xzvf - -C /opt

RUN mv /opt/druid-$DRUID_VERSION /opt/druid

WORKDIR /opt/druid

# install extensions via maven
# install to tmp directory because adding them directly
# to extensions directory causes an error
RUN java -classpath "lib/*" \
  -Ddruid.extensions.directory=/tmp/extensions \
  -Ddruid.extensions.hadoopDependenciesDir=/tmp/hadoop-dependencies \
  io.druid.cli.Main tools pull-deps \
  --defaultVersion $DRUID_VERSION \
  --no-default-hadoop \
  -c io.druid.extensions:mysql-metadata-storage \
  -c io.druid.extensions.contrib:statsd-emitter

RUN mv /tmp/extensions/* extensions
RUN rm -rf /tmp

VOLUME ["/data/druid", "/opt/druid/config", "/opt/druid/specs"]

COPY bin bin
COPY entrypoint.sh entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
