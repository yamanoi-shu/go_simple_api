#!/bin/bash

set -e

cmd="$@"

until mysqladmin ping -h db --silent; do
    echo 'waiting for mysqld to be connectable...'
    sleep 2
done

echo "go app is started!"

exec $cmd
