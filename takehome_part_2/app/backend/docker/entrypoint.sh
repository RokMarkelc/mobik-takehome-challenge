#!/bin/bash
set -e

migrate_version="v4.19.1"

echo "Installing migrate tool..."
curl -L https://github.com/golang-migrate/migrate/releases/download/$migrate_version/migrate.linux-amd64.tar.gz | tar xvz

echo "Running database migrations..."
./migrate -path=migrations -database postgres://$DB_USER:$DB_PASSWORD@$POSTGRES_SERVICE_HOST/app?sslmode=disable up
echo "Database migrations completed."

exec "$@"