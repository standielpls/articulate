.PHONY: all

VERSION=1.0.0
APP_NAME=articulate
PORT=8080
HOSTNAME=gcr.io/lancelot-274021

build:
	@go build -o bin/art cmd/main.go

deps:
	@go mod tidy
	@go mod vendor

docker-build:
	@docker build -t $(APP_NAME)  .

docker-run: 
	@docker run -it -e ADDR=:8080 -p 8080:8080 --rm $(APP_NAME)

docker-push:
	@gcloud auth configure-docker
	@docker tag $(APP_NAME) $(HOSTNAME)/$(APP_NAME):$(VERSION)
	@docker tag $(APP_NAME) $(HOSTNAME)/$(APP_NAME)
	@docker push $(HOSTNAME)/$(APP_NAME):$(VERSION)
	@docker push $(HOSTNAME)/$(APP_NAME)

env:
	@$(info    VERSION is $(VERSION))
	@$(info    APP_NAME is $(APP_NAME))
	@$(info    PORT is $(PORT))
	@$(info    HOSTNAME is $(HOSTNAME))

test:
	@go test ./...