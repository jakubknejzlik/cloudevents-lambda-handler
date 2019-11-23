package encoding

import "encoding/json"

type EventBridgeMessage struct {
	DetailType string      `json:"detail-type"` // In case of EventBridge, this value is nonempty string
	Detail     interface{} `json:"detail"`      // Message object
}

func DecodeEventBridgeMessage(body []byte) (message *EventBridgeMessage, ok bool) {
	ok = false
	err := json.Unmarshal(body, &message)
	if err == nil && message.IsValid() {
		ok = true
	}
	return
}

func (m *EventBridgeMessage) IsValid() bool {
	return m.DetailType != "" && m.Detail != nil
}

func (m *EventBridgeMessage) GetBody() []byte {
	body, _ := json.Marshal(m.Detail)
	return body
}
