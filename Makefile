generate-mocks:
	mockery --dir src/infra/repository/ --all --output src/infra/mocks
	mockery --dir src/infra/message_broker/ --all --output src/infra/mocks
	mockery --dir src/core/service/ --all --output src/core/mocks

services-down:
	docker compose down

terraform:
	terraform init
	terraform plan
	terraform apply --auto-approve

deploy-api:
	cd src/api && GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	cd src/api && zip -r ../../function.zip bootstrap . ../../go.mod ../../go.sum
	aws s3 cp function.zip s3://orders-service-01/function.zip --region us-east-1
	aws lambda update-function-code --function-name orders-service-lambda --s3-bucket orders-service-01 --s3-key function.zip --region us-east-1

deploy-async:
	cd src/async && GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
	cd src/async && zip -r ../../function.zip bootstrap . ../../go.mod ../../go.sum
	aws s3 cp function.zip s3://orders-service-01/function.zip --region us-east-1
	aws lambda update-function-code --function-name orders-service-async-lambda --s3-bucket orders-service-01 --s3-key function.zip --region us-east-1
