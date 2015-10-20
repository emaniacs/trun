GO=go
OUTPUT=trun-cli

main:
	$(GO) build -o $(OUTPUT)
clean:
	rm -f $(OUTPUT)
