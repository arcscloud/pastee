prod:
		GOOS=linux GOARCH=arm64 CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc go build -o server && docker build . --platform linux/arm64/v8 --tag arcscloud/pastee