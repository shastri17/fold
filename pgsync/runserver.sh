#! /bin/sh

./wait-for-it.sh $PG_HOST:$PG_PORT -t 60

./wait-for-it.sh $ELASTICSEARCH_HOST:$ELASTICSEARCH_PORT -t 60

./wait-for-it.sh $REDIS_HOST:$REDIS_PORT -t 60

EXAMPLE_DIR="pgsync"

bootstrap --config $EXAMPLE_DIR/schema.json

pgsync --config $EXAMPLE_DIR/schema.json --daemon