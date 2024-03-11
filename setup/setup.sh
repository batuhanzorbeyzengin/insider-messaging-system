#!/bin/bash

# Create a new Docker network
docker network create insider-network

# Run the MySQL container and connect it to the network
docker run --name insider-db --network insider-network -e MYSQL_ROOT_PASSWORD=root -d mysql:8

# Wait for the MySQL container to be ready
echo "Waiting for MySQL container to be ready..."
while ! docker exec insider-db mysqladmin ping --silent; do
    sleep 1
done

sleep 5

# Import the setup.sql file
echo "Importing setup.sql..."
docker exec -i insider-db mysql -h insider-db -P 3306 -uroot -proot < setup/setup.sql

# Run Docker Compose
echo "Starting Docker Compose..."
docker-compose up -d --build