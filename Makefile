GO=go
OUTPUT=trun
CLI_DIR=cli

main:
	cd $(CLI_DIR) ; $(GO) build -o ../$(OUTPUT)
test:
	$(GO) test
bench:
	$(GO) test -bench=.
clean:
	rm -f $(OUTPUT)
