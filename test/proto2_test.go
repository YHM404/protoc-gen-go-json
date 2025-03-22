package test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestProto2JsonSerialization(t *testing.T) {
	// 创建一个测试消息
	msg := &Proto2Message{
		RequiredField: proto.String("required value"),
		OptionalField: stringPtr("optional value"),
		RepeatedField: []int32{1, 2, 3},
		Nested: &Proto2Message_NestedMessage{
			InnerField: stringPtr("nested value"),
		},
		EnumField: Proto2Message_VALUE2.Enum(),
	}

	// 测试 JSON 序列化
	data, err := json.Marshal(msg)
	assert.NoError(t, err)

	// 验证序列化结果包含预期字段
	jsonString := string(data)
	assert.Contains(t, jsonString, `"requiredField":`)
	assert.Contains(t, jsonString, `"optionalField":`)
	assert.Contains(t, jsonString, `"repeatedField":`)
	assert.Contains(t, jsonString, `"nested":`)
	assert.Contains(t, jsonString, `"enumField":`)

	// 测试 JSON 反序列化
	var decodedMsg Proto2Message
	err = json.Unmarshal(data, &decodedMsg)
	assert.NoError(t, err)

	// 验证反序列化的消息与原消息相等
	assert.Equal(t, msg.RequiredField, decodedMsg.RequiredField)
	assert.Equal(t, *msg.OptionalField, *decodedMsg.OptionalField)
	assert.Equal(t, msg.RepeatedField, decodedMsg.RepeatedField)
	assert.Equal(t, *msg.Nested.InnerField, *decodedMsg.Nested.InnerField)
	assert.Equal(t, *msg.EnumField, *decodedMsg.EnumField)
}

func stringPtr(s string) *string {
	return &s
}
