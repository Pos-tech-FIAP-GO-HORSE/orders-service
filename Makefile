generate-mocks:
	mockery --dir src/infra/repository/ --all --output src/infra/mocks
	mockery --dir src/infra/message_broker/ --all --output src/infra/mocks
	mockery --dir src/core/service/ --all --output src/core/mocks

services-up:
	docker compose up -d
	terraform init
	terraform apply -auto-approve

services-down:
	docker compose down
