build:
	mkdir -p output
	go build -o output/bills cmd/cli/main.go
install:
	chmod +x output/bills
	sudo cp output/bills /usr/local/bin