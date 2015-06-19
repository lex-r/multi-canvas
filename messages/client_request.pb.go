// Code generated by protoc-gen-go.
// source: client_request.proto
// DO NOT EDIT!

package messages

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ClientRequest struct {
	Method             *string                   `protobuf:"bytes,1,req,name=method" json:"method,omitempty"`
	RequestPong        *ClientRequestPong        `protobuf:"bytes,2,opt,name=requestPong" json:"requestPong,omitempty"`
	RequestCreateWorld *ClientRequestCreateWorld `protobuf:"bytes,3,opt,name=requestCreateWorld" json:"requestCreateWorld,omitempty"`
	RequestMonitor     *ClientRequestMonitor     `protobuf:"bytes,4,opt,name=requestMonitor" json:"requestMonitor,omitempty"`
	XXX_unrecognized   []byte                    `json:"-"`
}

func (m *ClientRequest) Reset()         { *m = ClientRequest{} }
func (m *ClientRequest) String() string { return proto.CompactTextString(m) }
func (*ClientRequest) ProtoMessage()    {}

func (m *ClientRequest) GetMethod() string {
	if m != nil && m.Method != nil {
		return *m.Method
	}
	return ""
}

func (m *ClientRequest) GetRequestPong() *ClientRequestPong {
	if m != nil {
		return m.RequestPong
	}
	return nil
}

func (m *ClientRequest) GetRequestCreateWorld() *ClientRequestCreateWorld {
	if m != nil {
		return m.RequestCreateWorld
	}
	return nil
}

func (m *ClientRequest) GetRequestMonitor() *ClientRequestMonitor {
	if m != nil {
		return m.RequestMonitor
	}
	return nil
}

func init() {
}
