include .env

stop_container:
	@echo "Stopping container..."
	if [ $$(docker ps -q) ]; then \
  		echo "Found and Stopping container..."; \
		docker stop $$(docker ps -q); \
	else \
	  		echo "No container running..."; \
	fi


create_container:
	@echo "Creating container..."
	docker run --name ${DB_DOCKER_CONTAINER_NAME} -p 5433:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres:16.2-alpine

create_db:
	@echo "Creating database..."
	docker exec -it ${DB_DOCKER_CONTAINER_NAME} createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${DB_NAME}

start_container:
	@echo "Starting container..."
	docker start ${DB_DOCKER_CONTAINER_NAME}

create_migration:
	@echo "Creating migration..."
	sqlx migrate add -r init