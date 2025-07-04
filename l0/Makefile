DB_NAME = wb_orders
DB_USER = postgres

all:
	docker-compose build && docker-compose up -d 

down:
	docker-compose down -v

check_db_initial:
	docker exec orders-db psql -U $(DB_USER) -d $(DB_NAME) -c "SELECT * FROM orders;"
