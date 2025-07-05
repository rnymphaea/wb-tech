DB_NAME = wb_orders
DB_USER = postgres
APP_CONTAINER = orders-app
DAEMON_FLAG = -d

ifdef NO_DAEMON
    DAEMON_FLAG = 
endif

.PHONY: all
all: build up

.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up $(DAEMON_FLAG)

.PHONY: down
down:
	docker-compose down -v

.PHONY: restart
restart: down up

.PHONY: logs
logs:
	docker logs $(APP_CONTAINER)

.PHONY: check_db_initial
check_db_initial:
	docker exec orders-db psql -U $(DB_USER) -d $(DB_NAME) -c "SELECT * FROM orders;"

.PHONY: send_message
send_message:
	jq -c . order.json | docker exec -i kafka bash -c "kafka-console-producer --topic orders --bootstrap-server localhost:9093"
