package handler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type EventBridgeEvent struct {
	events.CloudWatchEvent `json:",inline"`
}

func DecodeEventBridgeEvent(body []byte) (message *EventBridgeEvent, ok bool) {
	ok = false
	err := json.Unmarshal(body, &message)
	if err == nil && message.IsValid() {
		ok = true
	}
	return
}

func (m *EventBridgeEvent) IsValid() bool {
	return m.DetailType != "" && m.Detail != nil
}

func (m *EventBridgeEvent) GetBodies() [][]byte {
	body, _ := json.Marshal(m.Detail)
	return [][]byte{body}
}
