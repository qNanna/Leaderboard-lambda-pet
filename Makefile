# build:
# 	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"

# zip:
# 	zip -FS ./.bin/bootstrap.zip ./.bin/bootstrap

npm.install.local:
	npm i --save-dev serverless serverless-offline serverless-go-plugin

npm.install.global:
	npm i -g serverless serverless-offline serverless-go-plugin

build:
	serverless go build

deploy:
	serverless deploy --force

remove:
	serverless remove

run:
	serverless offline --param='runtime=go1.x'

test:
	serverless invoke -f health 
