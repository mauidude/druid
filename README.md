# Event Aggregator

## Setup

1. [Create a new docker-machine](https://docs.docker.com/machine/get-started/).

```
# using virtual box
docker-machine create --driver virtualbox [name]

# set docker-machine env vars
eval $(docker-machine env [name])
```

*Note: the name you choose should be replaced in `bin/common.sh` as the `$DOCKER_MACHINE_NAME` variable*. The default name is assumed to be `dev`.

2. Bring up cluster

```
docker-compose up druid_broker druid_overlord druid_middle_manager
```

This will bring up all the dependencies of these services.

3. Query

You can now query the cluster once it's up and running. An example query is posted below. Please replace or set the `$DOCKER_MACHINE_IP` variable with the IP of your
docker-machine IP address. Queries are issued through the broker node (port 8082).

```
curl -XPOST -H 'Content-type: application/json' "http://$DOCKER_MACHINE_IP:8082/druid/v2/?pretty" -d '{
  "queryType": "groupBy",
  "dataSource": "sendgrid",
  "granularity": "second",
  "dimensions": ["type"],
  "aggregations": [
    { "type": "longSum", "name": "count", "fieldName": "count" }
  ],
  "intervals": [ "2012-01-01T00:00:00.000/2017-01-03T00:00:00.000" ]
}'
```

## Adding Test Data

When the cluster is up you can run the `event_pump` service to add data at one-second intervals.

```
docker-compose up event_pump
```


## Cluster Management

You can visit the management UI in your browser by going to:

```
http://192.168.99.100:8081/
```

The IP is the IP of your docker-machine which can be found by running `docker-machine env <name>`. The port should be the druid coordinator's port which is defaulted to 8081.

## Configuration

### Druid

These configurations are mostly the default ones which are NOT meant for production. Look at the documentation at druid.io to find appropriate production-level configurations.

Configurations are provided to Druid through property files located at [druid/config](druid/config). Environment variables will be substituted by using the syntax `$MY_ENV_VAR` in the file. The container(s) will need to be rebuilt after a change to the configs.

## Useful Documentation

- [Druid Overview](http://druid.io/docs/latest/design/design.html)


