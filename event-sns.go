package handler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type SNSEvent struct {
	events.SNSEvent `json:",inline"`
}

func DecodeSNSEvent(body []byte) (message *SNSEvent, ok bool) {
	ok = false
	err := json.Unmarshal(body, &message)
	if err == nil && message.IsNotification() {
		ok = true
	}
	return
}

func (m *SNSEvent) IsNotification() bool {
	return len(m.Records) > 0 && m.Records[0].EventSource == "aws:sns"
}

func (m *SNSEvent) GetBodies() [][]byte {
	res := [][]byte{}
	for _, r := range m.Records {
		res = append(res, []byte(r.SNS.Message))
	}
	return res
}
