#!/bin/sh

set -e

# for each env var, go in reverse order
# to prevent $ABC substituting for $ABC_XYZ
for VAR in `env | sort -r`
do
  # grab only variable name
  VAR=`echo "$VAR" | sed -r "s/^(.+?)=.*/\1/g"`

  # for each file containing that variable name
  grep -r -l "\$$VAR" config | while read -r FILE
  do
    # append $ to get variable name
    VALUE=$(echo "\$$VAR")
    VALUE=$(eval echo $VALUE)
    echo "setting $VAR to $VALUE in $FILE"

    sed -i "s/\$$VAR/$VALUE/g" $FILE
  done
done

exec "$@"
