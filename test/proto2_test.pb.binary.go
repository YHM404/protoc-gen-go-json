// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: proto2_test.proto

package test

import (
	"google.golang.org/protobuf/proto"
)

// Marshal marshals the message using proto.Marshal
func (msg *Proto2Message) Marshal() ([]byte, error) {
	return proto.Marshal(msg)
}

// Marshal marshals the message using proto.Marshal
func (msg *Proto2Message_NestedMessage) Marshal() ([]byte, error) {
	return proto.Marshal(msg)
}
