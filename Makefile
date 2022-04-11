prod:
		GOOS=linux GOARCH=arm64 go build -o server && docker build . --platform linux/arm64/v8 --tag arcscloud/pastee