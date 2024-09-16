agent:
	@echo "Building the crypto-agent"
	go build -C ./crypto-tracking-agent -v
	@echo "Done agent"
noti:
	@echo "Building the crypto-noti"
	go build -C ./crypto-tracking-notification -v
	@echo "Done noti"