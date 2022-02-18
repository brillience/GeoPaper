env-up:
	docker-compose -f docker-compose-env.yaml up -d
fill-db:
	go test -v ./internal/util/
	echo "Finished filling Database!"
run:
	go run main.go
env-down:
	docker-compose -f docker-compose-env.yaml down