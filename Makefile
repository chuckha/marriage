.PHONY: all
all: clean build upload

.PHONY: build
build: main deployment.zip

deployment.zip: main
	zip deployment.zip main

main:
	GOOS=linux go build -o main main.go

.PHONY: upload
upload:
	aws s3 cp deployment.zip s3://cupidsarrow/deployment.zip

.PHONY: clean
clean:
	rm deployment.zip main
