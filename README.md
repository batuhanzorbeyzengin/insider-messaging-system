# Insider Messaging System

This project is an implementation of an automatic message sending system built with Go and Fiber. It retrieves unsent messages from a MySQL database, sends them via a webhook, and caches the sent messages in Redis (bonus feature).

## Features

- Automatically sends unsent messages from the database every 2 minutes
- Supports character limit for message content
- Prevents resending messages that have already been sent
- Caches sent message IDs and timestamps in Redis (bonus feature)
- Provides two API endpoints:
  - Start/Stop automatic message sending
  - Retrieve a list of sent messages

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Fiber](https://gofiber.io/) - Web framework for Go
- [MySQL](https://www.mysql.com/) - Relational database for storing messages
- [Redis](https://redis.io/) - In-memory data store for caching sent messages
- [Swagger](https://swagger.io/) - API documentation tool

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.21 or later)
- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository: git clone https://github.com/batuhanzorbeyzengin/insider-messaging-system
2. Navigate to the project directory: cd insider-messaging-system
3. Copy the `.env.example` file to a new file named `.env`: `cp .env.example .env`
4. Open the `.env` file and fill in the required environment variables with your specific values.
5. Run the `setup.sh` script to create the required Docker network, start the MySQL container, and run Docker Compose: `./setup/setup.sh`
This script will create a Docker network named "insider-network", start the MySQL container, import the initial database setup, and then start the Docker Compose services.
6. Access the application at `http://localhost:3000` (or the port specified in the `SERVER_PORT` environment variable).

7. You can follow the messages sent by the application at [`https://webhook.site/2baa69ce-0241-4a7b-8fc9-4845a3206e9c`](`https://webhook.site/2baa69ce-0241-4a7b-8fc9-4845a3206e9c`).
### API Documentation

The API documentation is available at `http://localhost:3000/swagger/index.html` (or the appropriate URL based on your `SERVER_PORT` setting). You can explore and test the available endpoints using the Swagger UI.

### Development

If you want to run the application locally without Docker, follow these steps:

1. Install and set up MySQL and Redis on your local machine.
2. Update the `.env` file with the correct database and Redis connection details.
3. Run the database migrations: go run migrations/migration.go
4. Start the Go application: go run cmd/main.go

The application will be accessible at `http://localhost:3000` (or the specified `SERVER_PORT`).