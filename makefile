build-example:
	GO111MODULE=on GOOS=linux go build -o main example/handler.go && zip lambda.zip main && rm main