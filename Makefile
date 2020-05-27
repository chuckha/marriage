deployment.zip: main
	zip deployment.zip main

main:
	GOOS=linux go build -o main main.go

.PHONY: clean
clean:
	rm deployment.zip main
