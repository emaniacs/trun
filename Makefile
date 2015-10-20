GO=go
OUTPUT=../trun
CLI_DIR=cli

main:
	cd $(CLI_DIR) ; $(GO) build -o $(OUTPUT)
clean:
	rm -f $(OUTPUT)
