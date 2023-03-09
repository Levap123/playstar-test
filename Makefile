build:
	docker-compose build

run:	
	docker-compose up
	docker exec rabbitmq rabbitmqadmin declare queue name=logs vhost=my_vhost
