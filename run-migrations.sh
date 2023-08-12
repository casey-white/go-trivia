#!/bin/sh

echo "Running migrations..."

for file in /migrations/*.sql
do
    echo "Running $file..."
    psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$file"
done