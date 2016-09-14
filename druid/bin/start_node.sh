#!/usr/bin/env bash
set +e
set +u
shopt -s xpg_echo
shopt -s expand_aliases
trap "exit 1" 1 2 3 15

SCRIPT_DIR=$( cd $(dirname $0)/.. && pwd )
EXTENSIONS_DIR="${SCRIPT_DIR}/extensions"
SERVER_TYPE="$1"

if [ -z "${SERVER_TYPE// }" ]
then
  echo "usage: $0 server-type" >& 2
  exit 2
fi

if [[ -z "${JAVA_MX// }" ]]
then
  JAVA_MX="512m"
  echo "defaulting to -Xmx$JAVA_MX"
fi

if [[ ! -d "${SCRIPT_DIR}/lib" || ! -d "${SCRIPT_DIR}/config" ]]; then
  echo "This script appears to be running from the source location. It must be run from its deployed location."
  echo "After building, unpack services/target/druid-services-*-SNAPSHOT-bin.tar.gz, and run the script unpacked there."
  exit 2
fi

CURR_DIR=`pwd`
cd ${SCRIPT_DIR}
SCRIPT_DIR=`pwd`
cd ${CURR_DIR}

# start process
JAVA_ARGS="${JAVA_ARGS} -Xmx${JAVA_MX} -Duser.timezone=UTC -Dfile.encoding=UTF-8"
JAVA_ARGS="${JAVA_ARGS} -Ddruid.extensions.directory=${EXTENSIONS_DIR}"

DRUID_CP="${SCRIPT_DIR}/config/_common"
DRUID_CP="${DRUID_CP}:${SCRIPT_DIR}/config/$SERVER_TYPE"
DRUID_CP="${DRUID_CP}:${SCRIPT_DIR}/lib/*"
DRUID_CP="${DRUID_CP}:${EXTENSIONS_DIR}/*"

echo "using classpath $DRUID_CP in `pwd`"
echo "using java args $JAVA_ARGS"

exec java ${JAVA_ARGS} -classpath "${DRUID_CP}" io.druid.cli.Main server "$SERVER_TYPE"
