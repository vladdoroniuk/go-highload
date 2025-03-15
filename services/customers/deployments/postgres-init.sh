#!/bin/bash

echo "Waiting for PostgreSQL to be ready..."

until psql -h go-highload-postgres -U "$POSTGRES_USER" -d postgres -c "SELECT 1" &>/dev/null; do
    sleep 5
    echo "Still waiting for PostgreSQL..."
done

echo "PostgreSQL is ready! Creating database if not exists..."
psql -h go-highload-postgres -U "$POSTGRES_USER" -d postgres -c "CREATE DATABASE $POSTGRES_DB;" || echo "Database $POSTGRES_DB already exists."
echo "Granting privileges..."
psql -h go-highload-postgres -U "$POSTGRES_USER" -d postgres -c "GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;"
echo "Done!"