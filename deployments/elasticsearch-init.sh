#!/bin/sh

echo "Waiting for Elasticsearch to be ready..."

until curl -s -u "${ELASTICSEARCH_USER}:${ELASTICSEARCH_PASSWORD}" "http://go-highload-elasticsearch:${ELASTICSEARCH_PORT}" >/dev/null; do
    sleep 5
    echo "Still waiting for Elasticsearch..."
done

echo "Elasticsearch is ready! Creating index '${ELASTICSEARCH_INDEX}'..."

curl -X PUT "http://go-highload-elasticsearch:${ELASTICSEARCH_PORT}/${ELASTICSEARCH_INDEX}" \
     -H "Content-Type: application/json" \
     -u "${ELASTICSEARCH_USER}:${ELASTICSEARCH_PASSWORD}" \
     -d '{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 0
  }
}'

echo "Done!"