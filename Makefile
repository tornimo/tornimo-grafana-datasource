all: webpack build

webpack:
	webpack

build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/tornimo-datasource-plugin_linux_amd64 ./pkg
	GOOS=darwin GOARCH=amd64 go build -o ./dist/tornimo-datasource-plugin_darwin_amd64 ./pkg
