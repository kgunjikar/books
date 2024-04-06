all:
	go mod tidy
	cp html/* /tmp/data
	go build main.go
clean:
	rm -f main
