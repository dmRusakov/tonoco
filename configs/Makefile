SERVICE_PATH=`basename $(CURDIR)`

up:
	docker-compose --env-file ./dev.env up -d
	@echo "Started"

down:
	docker-compose --env-file ./dev.env down --volumes
	@echo "Stopped"

reset: down delete up
	@echo "Reseted"

restart:
	docker-compose --env-file ./dev.env restart
	@echo "Restarted"

delete:
	docker-compose --env-file ./dev.env down --volumes
	cd ../${SERVICE_PATH}/dbData/DataStorage/ \
	&& sudo chown -R 1000:1000 data/ \
	&& rm -rf data/
	@echo "Deleted"