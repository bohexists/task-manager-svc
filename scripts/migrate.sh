#!/bin/bash

set -e

DB_HOST=${DB_HOST:-"localhost"}
DB_PORT=${DB_PORT:-"3306"}
DB_USER=${DB_USER:-"root"}
DB_PASSWORD=${DB_PASSWORD:-"secret"}
DB_NAME=${DB_NAME:-"task_db"}

# Path to migrations
MIGRATIONS_DIR="./migrations"

# Migration command
MIGRATE_CMD="migrate -path $MIGRATIONS_DIR -database mysql://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME"

function migrate_up() {
    echo "Applying migrations..."
    $MIGRATE_CMD up
    echo "Migrations applied successfully."
}

function migrate_down() {
    echo "Rolling back the last migration..."
    $MIGRATE_CMD down 1
    echo "Rolled back successfully."
}

function migrate_status() {
    echo "Checking migration status..."
    $MIGRATE_CMD version
}

function migrate_force() {
    echo "Forcing migration version..."
    if [ -z "$1" ]; then
        echo "You must provide a version number to force."
        exit 1
    fi
    $MIGRATE_CMD force $1
}

# Check command
case "$1" in
    "up")
        migrate_up
        ;;
    "down")
        migrate_down
        ;;
    "status")
        migrate_status
        ;;
    "force")
        migrate_force $2
        ;;
    *)
        echo "Usage: $0 {up|down|status|force <version>}"
        exit 1
        ;;
esac