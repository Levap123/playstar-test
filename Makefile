build:
	docker-compose build

run:	
	docker-compose up
	migrate -path migrations/ -database "postgresql://root:root@localhost:5432/logs?sslmode=disable" -verbose up
	docker exec rabbitmq rabbitmqadmin declare queue name=logs vhost=my_vhost
