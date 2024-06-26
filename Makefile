-include .env

all: clean install test build

.PHONY: dev
dev : dev_one

.PHONY: dev_one
dev_one:; MNEMONIC="infant apart enroll relief kangaroo patch awesome wagon trap feature armor approve" go run . --yaml ./defaults/config.local.yml

.PHONY: dev_two
dev_two:; MNEMONIC="shy smile praise educate custom fashion gun enjoy zero powder garden second" go run . --yaml ./defaults/config.local.yml

.PHONY: dev_three
dev_three:; MNEMONIC="wink giant track dwarf visa feed visual drip play grant royal noise" go run . --yaml ./defaults/config.local.yml

.PHONY: clean
clean: clean_tmp_data
	go clean && go mod tidy

.PHONY: clean_tmp_data
clean_tmp_data :; if [ -d "/tmp/data" ]; then sudo rm -rf /tmp/data; fi

.PHONY: install
install :; go mod download && go mod verify

.PHONY: lint
lint :; golangci-lint run

.PHONY: test
test :; go test -v ./...

.PHONY: test_coverage
test_coverage :; bash ./coverage.sh 

.PHONY: test_coverage_html
open_test_coverage :; bash ./coverage.sh && open ./coverage.html

.PHONY: build
build :; go build -o wpokt-oracle .

.PHONY: docker_build
docker_build :; docker buildx build . -t dan13ram/wpokt-oracle:v0.0.1 --file ./docker/Dockerfile

.PHONY: docker_push
docker_push :; docker push dan13ram/wpokt-oracle:v0.0.1

.PHONY: docker_dev
docker_dev : docker_one

.PHONY: docker_one
docker_one :; MNEMONIC="infant apart enroll relief kangaroo patch awesome wagon trap feature armor approve" YAML_FILE=/app/defaults/config.local.yml docker compose -f docker/docker-compose.yml up --force-recreate

.PHONY: docker_two
docker_two :; MNEMONIC="shy smile praise educate custom fashion gun enjoy zero powder garden second" YAML_FILE=/app/defaults/config.local.yml docker compose -f docker/docker-compose.yml up --force-recreate

.PHONY: docker_three
docker_three :; MNEMONIC="wink giant track dwarf visa feed visual drip play grant royal noise" YAML_FILE=/app/defaults/config.local.yml docker compose -f docker/docker-compose.yml up --force-recreate

.PHONY: localnet_up
localnet_up:; docker compose -f e2e/docker-compose.yml up --force-recreate

.PHONY: prompt_user
prompt_user:
	@echo "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

.PHONY: docker_wipe
docker_wipe: prompt_user ## [WARNING] Remove all the docker containers, images and volumes.
	docker ps -a -q | xargs -r -I {} docker stop {}
	docker ps -a -q | xargs -r -I {} docker rm {}
	docker images -q | xargs -r -I {} docker rmi {}
	docker volume ls -q | xargs -r -I {} docker volume rm {}

.PHONY: e2e_test
e2e_test :; cd e2e && yarn install && yarn test

.PHONY: generate_keys
generate_keys :; go run scripts/generate_keys/main.go --mnemonic "${mnemonic}"

.PHONY: generate_multisig
generate_multisig :; go run scripts/generate_multisig/main.go --publickeys "${publickeys}" --threshold ${threshold}
