# go build command
gobuild:
	@go build -v -o team6minihack cmd/team6minihack/*.go

# go run command
gorun:
	make gobuild
	@./team6minihack --config_path='./config/tkp-app.{TKPENV}.yaml'

# deploy command
deploy:
	@echo "deploying"