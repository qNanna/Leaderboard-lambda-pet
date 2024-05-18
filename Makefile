# -ldflags="-s -w" (stripped -25% binary without debug) # -gcflags "all=-N -l" (debug, but m2 still have error)
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
	node node_modules/serverless/bin/serverless deploy --force

remove:
	node node_modules/serverless/bin/serverless remove

run:
	node node_modules/serverless/bin/serverless offline --param='runtime=go1.x'

test:
	node node_modules/serverless/bin/serverless invoke -f health 
