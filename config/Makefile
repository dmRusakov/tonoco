CONFIG_FOLDER=`basename $(CURDIR)`

admin:

up: down
	docker-compose --env-file ./dev.env up -d
	@echo "Started"

down:
	docker-compose --env-file ./dev.env down --volumes
	@echo "Stopped"

reset: down delete up down up
	@echo "Reseted"

restart:
	docker-compose --env-file ./dev.env restart
	@echo "Restarted"

delete:
	docker-compose --env-file ./dev.env down --volumes
	sudo ./deleteDbData.sh
	sleep 1 # wait 1 sec for db to be deleted
	@echo "Deleted"

generate_grpc:
	@$(GOPATH)/bin/buf generate
	@for dir in $(CURDIR)/$$dir/gen/go/*/; do \
	  cd $$dir && \
	  go mod init && go mod tidy; \
	done