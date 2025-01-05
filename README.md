
# Go Items Management with Queues

## Overview

This project demonstrates the use of **Golang microservices** to manage items in a store. It showcases the implementation of a producer-consumer architecture, leveraging **RabbitMQ** for asynchronous communication, **PostgreSQL** for data storage, and **Redis** for caching.

### Key Features

- **Hexagonal Architecture**: Implements ports and adapters, adhering to Domain-Driven Design (DDD) principles.
- **Microservices**:
  - **API Service**: Registers items into a queue.
  - **Worker (Consumer)**: Processes the queue, updates the database, and generates a Redis cache.
- **CRUD Operations**: Supports creating, reading, updating, and deleting items.
- **Asynchronous Communication**: All operations are processed via RabbitMQ queues, emphasizing decoupled components.

> **Note**: While asynchronous CRUD operations might not be optimal for typical use cases, this project aims to demonstrate specific tools and design patterns.

---

## Getting Started

### Prerequisites

Ensure the following are installed on your system:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- Golang (optional, if running locally)

### Environment Variables

This project haves a `.environment.yaml` file in the root directory with the following variables:

```yaml
database:
  user: postgres
  password: postgres
  host: database
  port: 5432
  dbname: management
  driver: postgres

rabbitmq:
  user: rabbitmq
  password: rabbitmq
  address: rabbitmq
  port: 5672
  vhost: /
  retry-policy:
    max-retries: 3
    interval: 5

redis:
  address: redis
  port: 6379
  password: redis
```

---

## How to Run

To start the services, navigate to the project directory and run:

```bash
docker compose -f build/docker-compose.yaml up -d
```
or (if you have make instaled):
```bash
make up
```

To stop the services, use:

```bash
docker compose -f build/docker-compose.yaml down
```
or (if you have make instaled):
```bash
make down
```

### Logs

- API Service Logs:
  ```bash
  docker logs -f items-management-api
  ```
  or (if you have make instaled):
  ```bash
  make logs-api
  ```
- Consumer Service Logs:
  ```bash
  docker logs -f items-management-consumer
  ```
  or (if you have make instaled):
  ```bash
  make logs-consumer
  ```

---

## Usage

### API Endpoints

- **POST /items**: Create a new item.

#### Example Request

```http
POST http://localhost:8080/items
Content-Type: application/json

{
    "name": "Socks",
    "description": "Wonderful socks",
    "price": 27.99,
    "stock": 10,
    "status": "available"
}
```

### Docker Services Overview

| Service   | Description                             | Ports         |
|-----------|-----------------------------------------|---------------|
| `api`     | API service for managing items          | `8080:8080`   |
| `consumer`| Worker service for queue consumption    | -             |
| `rabbitmq`| RabbitMQ message broker                 | `5672`, `15672`|
| `database`| PostgreSQL database                     | `5432`        |
| `redis`   | Redis cache                             | `6379`        |
| `pgadmin` | PostgreSQL management interface         | `5050`        |

---

## Project Structure

- **API**: Handles incoming HTTP requests and queues messages.
- **Worker**: Processes messages from the queue, updates the database, and caches data.
- **Database**: Stores persistent data using PostgreSQL.
- **Cache**: Improves read performance with Redis.
- **Messaging**: Facilitates asynchronous communication using RabbitMQ.

---

## Development Notes

### Architecture

This project uses **Hexagonal Architecture**, with the following components:

- **Domain**: Core business logic.
- **Use Cases**: Application-specific operations.
- **Adapters**: Interfaces for external systems (e.g., database, message broker).
