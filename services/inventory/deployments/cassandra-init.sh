#!/bin/bash

echo "Waiting for Cassandra to be ready..."

until cqlsh go-highload-cassandra -e "DESCRIBE KEYSPACES" &>/dev/null; do
    sleep 5
    echo "Still waiting for Cassandra..."
done

echo "Cassandra is ready! Creating keyspace if not exists..."

cqlsh go-highload-cassandra -u "$CASSANDRA_USER" -p "$CASSANDRA_PASSWORD" -e "CREATE KEYSPACE IF NOT EXISTS $CASSANDRA_KEYSPACE WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};"

echo "Done!"