package main

import (
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go"
	handler "github.com/jakubknejzlik/cloudevents-lambda-handler"
)

func receiver(e cloudevents.Event) error {
	fmt.Println("Received event", e)
	return nil
}

func main() {
	h := handler.NewCloudEventsLambdaHandler(receiver)
	h.Start()
}
