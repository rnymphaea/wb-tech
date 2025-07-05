DB_NAME = wb_orders
DB_USER = postgres

all:
	docker-compose build && docker-compose up -d 

down:
	docker-compose down -v

.PHONY: check_db_initial
check_db_initial:
	docker exec orders-db psql -U $(DB_USER) -d $(DB_NAME) -c "SELECT * FROM orders;"

.PHONY: send_message
send_message:
	jq -c . order.json | docker exec -i kafka bash -c "kafka-console-producer --topic orders --bootstrap-server localhost:9093"
