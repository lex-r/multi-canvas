// Code generated by protoc-gen-go.
// source: client_request_pong.proto
// DO NOT EDIT!

package messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ClientRequestPong struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientRequestPong) Reset()         { *m = ClientRequestPong{} }
func (m *ClientRequestPong) String() string { return proto.CompactTextString(m) }
func (*ClientRequestPong) ProtoMessage()    {}

func (m *ClientRequestPong) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func init() {
}
