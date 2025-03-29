#!/bin/bash

echo "Waiting for PostgreSQL to be ready..."

until pg_isready -h go-highload-postgres -U "$POSTGRES_USER"; do
    sleep 5
    echo "Still waiting for PostgreSQL..."
done

echo "PostgreSQL is ready! Creating database if not exists..."

export PGPASSWORD="$POSTGRES_PASSWORD"
psql -h go-highload-postgres -U "$POSTGRES_USER" -d postgres -tc "SELECT 1 FROM pg_database WHERE datname = '$POSTGRES_DB'" | grep -q 1 || \
psql -h go-highload-postgres -U "$POSTGRES_USER" -d postgres -c "CREATE DATABASE $POSTGRES_DB;"

echo "Done!"