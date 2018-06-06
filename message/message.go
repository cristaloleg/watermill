package message

import (
	"github.com/mitchellh/mapstructure"
)

type Payload interface{}

type Message interface {
	UUID() string

	SetMetadata(key, value string)
	GetMetadata(key string) string

	UnmarshalPayload(val interface{}) error
}

type Default struct {
	MessageUUID     string            `json:"message_uuid"`
	MessageMetadata map[string]string `json:"message_metadata"`
	MessagePayload  Payload           `json:"message_payload"`
}

func NewDefault(uuid string, payload Payload) Message {
	return &Default{
		MessageUUID:     uuid,
		MessageMetadata: make(map[string]string),
		MessagePayload:  payload,
	}
}

func (m Default) UUID() string {
	return m.MessageUUID
}

func (m *Default) SetMetadata(key, value string) {
	m.MessageMetadata[key] = value
}

func (m *Default) GetMetadata(key string) string {
	if val, ok := m.MessageMetadata[key]; ok {
		return val
	}

	return ""
}

func (m *Default) UnmarshalPayload(val interface{}) error {
	return mapstructure.Decode(m.MessagePayload, val)
}