
# Order Tracking Service

## Overview

A microservice for tracking and viewing order details with the following features:
- Web interface for order lookup
- RESTful API endpoints
- PostgreSQL database storage
- Kafka integration for order processing
- Redis caching layer

## Project Structure
```
.
├── cmd
│   └── main.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── database
│   │   ├── cache
│   │   │   └── redis
│   │   │       └── redis.go
│   │   ├── migrations
│   │   │   ├── 000001_init_tables_up.sql
│   │   │   └── initial_script.sql
│   │   ├── models
│   │   │   └── models.go
│   │   ├── postgres
│   │   │   ├── orders.go
│   │   │   └── postgres.go
│   │   └── storage.go
│   ├── handler
│   │   ├── handler.go
│   │   ├── order.go
│   │   └── templates
│   │       ├── index.html
│   │       └── order.html
│   ├── kafka
│   │   └── consumer.go
│   └── router
│       └── router.go
├── Makefile
└── order.json
```

## Techologies

- Go
- Docker && docker-compose
- PostgreSQL
- Kafka
- Redis

## Getting Started

### 1. Clone the repository

```bash
git clone git@github.com:rnymphaea/wb-tech.git
cd wb-tech/l0
```
### 2. Setup env variables
### 3. Start service with make
```bash
make
```

## API Endpoints
### Web Interface
- `GET /` - Order lookup form
- `GET /order/{uid}` - Order details page

## Development
### make commands
#### Basic operations
- `make` - Build and start all containers (equivalent to `make build up`)
- `make build` - Build Docker containers
- `make up` - Start all containers in detached mode (add `NO_DAEMON=1` to run in foreground)
- `make down` - Stop and remove all containers with volumes
- `make restart` - Restart all containers (down then up)
- `make logs` - View logs from the application container

#### Database operations
- `make check_db_initial` - Verify database initialization by checking orders table

#### Kafka operations
- `make send_message` - Send a test message to Kafka using order.json

### order.json example
Example for kafka: `make send_message`
```
{
   "order_uid": "b563feb7b2b84b6tost",
   "track_number": "WBILMTESTTRACK",
   "entry": "WBIL",
   "delivery": {
      "name": "Test Testov",
      "phone": "+9720000000",
      "zip": "2639809",
      "city": "Kiryat Mozkin",
      "address": "Ploshad Mira 15",
      "region": "Kraiot",
      "email": "test@gmail.com"
   },
   "payment": {
      "transaction": "b563feb7b2b84b6tost",
      "request_id": "",
      "currency": "USD",
      "provider": "wbpay",
      "amount": 1817,
      "payment_dt": 1637907727,
      "bank": "alpha",
      "delivery_cost": 1500,
      "goods_total": 317,
      "custom_fee": 0
   },
   "items": [
      {
         "chrt_id": 9934930,
         "track_number": "WBILMTESTTRACK",
         "price": 453,
         "rid": "ab4219087a764ae0btost",
         "name": "Mascaras",
         "sale": 30,
         "size": "0",
         "total_price": 317,
         "nm_id": 2389212,
         "brand": "Vivienne Sabo",
         "status": 202
      }
   ],
   "locale": "en",
   "internal_signature": "",
   "customer_id": "test",
   "delivery_service": "meest",
   "shardkey": "9",
   "sm_id": 99,
   "date_created": "2021-11-26T06:22:19Z",
   "oof_shard": "1"
}
```

### .env example
```
DB_HOST=host
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=dbname

POSTGRES_PORT=5433

REDIS_ADDR=redis:6379
REDIS_TTL=86400s

KAFKA_BROKERS=kafka:9092
KAFKA_TOPIC=topic
KAFKA_GROUP_ID=group

ZOOKEEPER_CLIENT_PORT=2181

KAFKA_BROKER_ID=1
KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092,PLAINTEXT_HOST://localhost:9093

```
