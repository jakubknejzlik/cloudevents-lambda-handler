# cloudevents-lambda-handler

AWS Lambda handler for processing CloudEvents (from SNS/SQS/EventBridge).
Create CloudEvent receiver without the need to handle the AWS Event parsing/detection.

Currently supported event sources:

- `AWS EventBridge`
- `AWS SQS`
- `AWS SNS`

All received events are unwrapped and decoded to CloudEvent events.

For sending events checkout the [CloudEvent AWS transport](https://github.com/jakubknejzlik/cloudevents-aws-transport)

# Getting Started

Create your `main.go` file with this content:

```
package main

import (
	cloudevents "github.com/cloudevents/sdk-go"
	handler "github.com/jakubknejzlik/cloudevents-lambda-handler"
)

func receiver(e cloudevents.Event) error {
	// fmt.Println("Received event", e)
	return nil
}

func main() {
	h := handler.NewCloudEventsLambdaHandler(receiver)
	h.Start()
}

```

Build archive with AWS Lambda sources by running:

```
GOOS=linux go build -o main main.go && zip lambda.zip main && rm main
```
