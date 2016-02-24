#!/bin/sh

set -e

export KAFKA_BROKER_LIST="$KAFKA_PORT_9092_TCP_ADDR:$KAFKA_PORT_9092_TCP_PORT"
echo "using broker list: $KAFKA_BROKER_LIST"

exec "$@"
