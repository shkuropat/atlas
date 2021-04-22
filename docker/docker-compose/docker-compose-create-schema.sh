#!/bin/bash

KAFKA_HOST=kafka
KAFKA_PORT=9093
KAFKA_TOPIC=qwe
# Create kafka topic
docker-compose -f ./docker-compose.yaml run --entrypoint=/bin/sh kafka -c "kafka-topics.sh --bootstrap-server=${KAFKA_HOST}:${KAFKA_PORT} --create --topic ${KAFKA_TOPIC}"
docker-compose -f ./docker-compose.yaml run --entrypoint=/bin/sh kafka -c "kafka-topics.sh --bootstrap-server=${KAFKA_HOST}:${KAFKA_PORT} --list"

SCHEMA_FILE=/trail/clickhouse/journal_clickhouse_schema.sql
CLICKHOUSE_HOST=clickhouse-server
# Create ClickHouse schema
docker-compose -f ./docker-compose.yaml run --entrypoint=/bin/sh clickhouse-client -c "cat ${SCHEMA_FILE} | clickhouse-client --host=${CLICKHOUSE_HOST} --multiline --multiquery"

MINIO_HOST=minio
MINIO_PORT=9000
MINIO_ACCESS_KEY=minio1
MINIO_SECRET_KEY=minio123
BUCKET_NAME=mybucket
# Create MinIO bucket
docker-compose -f ./docker-compose.yaml run --entrypoint=/bin/sh minio-mc -c "mc config host add minio-docker http://${MINIO_HOST}:${MINIO_PORT} ${MINIO_ACCESS_KEY} ${MINIO_SECRET_KEY}; mc mb minio-docker/${BUCKET_NAME}; mc ls minio-docker"
