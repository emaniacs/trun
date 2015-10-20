GO=go
OUTPUT=trun-cli
CLI_DIR=cli

main:
	cd $(CLI_DIR)
	$(GO) build -o $(OUTPUT)
clean:
	rm -f $(OUTPUT)
