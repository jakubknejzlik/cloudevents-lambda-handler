package handler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type SQSEvent struct {
	events.SQSEvent `json:",inline"`
}

func DecodeSQSEvent(body []byte) (message *SQSEvent, ok bool) {
	ok = false
	err := json.Unmarshal(body, &message)
	if err == nil && message.IsValid() {
		ok = true
	}
	return
}

func (m *SQSEvent) IsValid() bool {
	return len(m.Records) > 0 && m.Records[0].EventSource == "aws:sqs"
}

func (m *SQSEvent) GetBodies() [][]byte {
	res := [][]byte{}
	for _, r := range m.Records {
		res = append(res, []byte(r.Body))
	}
	return res
}
