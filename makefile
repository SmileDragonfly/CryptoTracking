crypto-agent:
	@echo "Building the crypto-agent"
	go build -C ./crypto-tracking-agent -v
	@echo "Done"