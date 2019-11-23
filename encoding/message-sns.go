package encoding

import "encoding/json"

type SNSMessage struct {
	Type    string // In case of SNS, the type is "Notification"
	Message string // JSON encoded message
}

func DecodeSNSMessage(body []byte) (message *SNSMessage, ok bool) {
	ok = false
	err := json.Unmarshal(body, &message)
	if err == nil && message.IsNotification() {
		ok = true
	}
	return
}

func (m *SNSMessage) IsNotification() bool {
	return m.Type == "Notification"
}

func (m *SNSMessage) GetBody() []byte {
	return []byte(m.Message)
}
