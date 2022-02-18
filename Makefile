env-up:
	docker-compose -f docker-compose-env.yaml up -d
env-down:
	docker-compose -f docker-compose-env.yaml down
fill-db:
	go test -v ./internal/util/
	echo "Finished filling Database!"