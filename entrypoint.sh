#!/bin/sh
set -e

echo "Ожидание PostgreSQL на $POSTGRES_HOST:$POSTGRES_PORT..."
while ! nc -z $POSTGRES_HOST $POSTGRES_PORT; do
  sleep 0.1
done
echo "PostgreSQL доступен"

echo "Запуск миграций Goose..."
goose -dir migrations postgres \
  "user=$POSTGRES_USER \
  password=$POSTGRES_PASSWORD \
  dbname=$POSTGRES_DB \
  host=$POSTGRES_HOST \
  port=$POSTGRES_PORT \
  sslmode=disable" \
  up

echo "Запуск приложения..."
exec ./bank
