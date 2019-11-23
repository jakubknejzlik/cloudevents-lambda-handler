package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/cloudevents/sdk-go/pkg/cloudevents/transport"
	"github.com/jakubknejzlik/cloudevents-lambda-handler/encoding"
)

type CloudEventsReceiver = func(cloudevents.Event) error

type CloudEventsLambdaHandler struct {
	receiver CloudEventsReceiver
	codec    transport.Codec
}

func (h *CloudEventsLambdaHandler) Start() {
	lambda.StartHandler(h)
}

func (h *CloudEventsLambdaHandler) Invoke(ctx context.Context, payload []byte) (res []byte, err error) {
	var messageBodies [][]byte

	if m, ok := DecodeEventBridgeEvent(payload); ok {
		messageBodies = m.GetBodies()
	} else if m, ok := DecodeSNSEvent(payload); ok {
		messageBodies = m.GetBodies()
	} else if m, ok := DecodeSQSEvent(payload); ok {
		messageBodies = m.GetBodies()
	}

	for _, body := range messageBodies {
		msg := &encoding.Message{
			Body: body,
		}
		event, _err := h.codec.Decode(ctx, msg)
		if _err != nil {
			err = _err
			return
		}
		err = h.receiver(*event)
		if err != nil {
			return
		}
	}

	return
}

func NewCloudEventsLambdaHandler(receiver CloudEventsReceiver) CloudEventsLambdaHandler {
	codec := &encoding.Codec{}
	return CloudEventsLambdaHandler{receiver: receiver, codec: codec}
}
