[
  {
    "dataSchema": {
      "dataSource": "sendgrid",
      "parser": {
        "type": "string",
        "parseSpec": {
          "format": "json",
          "timestampSpec": {
            "column": "timestamp",
            "format": "ruby"
          },
          "dimensionsSpec": {
            "dimensions": ["event", "type", "ip", "sg_message_id"],
            "dimensionExclusions": [
              "sg_event_id",
              "smtp-id",
              "cert_err",
              "email",
              "useragent",
              "category",
              "pool",
              "reason",
              "newsletter",
              "response",
              "attempt"
            ],
            "spatialDimensions": []
          }
        }
      },
      "metricsSpec": [{
        "type": "count",
        "name": "count"
      }],
      "granularitySpec": {
        "type": "uniform",
        "segmentGranularity": "minute",
        "queryGranularity": "minute"
      }
    },
    "ioConfig": {
      "type": "realtime",
      "firehose": {
        "type": "kafka-0.8",
        "consumerProps": {
          "zookeeper.connect": "zookeeper:2181",
          "zookeeper.connection.timeout.ms": "15000",
          "zookeeper.session.timeout.ms": "15000",
          "zookeeper.sync.time.ms": "5000",
          "group.id": "druid-events",
          "fetch.message.max.bytes": "1048586",
          "auto.offset.reset": "largest",
          "auto.commit.enable": "false"
        },
        "feed": "events"
      },
      "plumber": {
        "type": "realtime"
      }
    },
    "tuningConfig": {
      "type": "realtime",
      "maxRowsInMemory": 500000,
      "intermediatePersistPeriod": "PT10m",
      "windowPeriod": "PT10m",
      "basePersistDirectory": "\/data\/druid/realtime\/basePersist",
      "rejectionPolicy": {
        "type": "messageTime"
      }
    }
  }
]
