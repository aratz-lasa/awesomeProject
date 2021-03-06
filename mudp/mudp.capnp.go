// Code generated by capnpc-go. DO NOT EDIT.

package main

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
	strconv "strconv"
)

type MudpRequest struct{ capnp.Struct }

// MudpRequest_TypeID is the unique identifier for the type MudpRequest.
const MudpRequest_TypeID = 0xfd0cdac648e8286b

func NewMudpRequest(s *capnp.Segment) (MudpRequest, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return MudpRequest{st}, err
}

func NewRootMudpRequest(s *capnp.Segment) (MudpRequest, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return MudpRequest{st}, err
}

func ReadRootMudpRequest(msg *capnp.Message) (MudpRequest, error) {
	root, err := msg.Root()
	return MudpRequest{root.Struct()}, err
}

func (s MudpRequest) String() string {
	str, _ := text.Marshal(0xfd0cdac648e8286b, s.Struct)
	return str
}

func (s MudpRequest) Peer() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s MudpRequest) HasPeer() bool {
	return s.Struct.HasPtr(0)
}

func (s MudpRequest) SetPeer(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s MudpRequest) Namespace() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s MudpRequest) HasNamespace() bool {
	return s.Struct.HasPtr(1)
}

func (s MudpRequest) NamespaceBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s MudpRequest) SetNamespace(v string) error {
	return s.Struct.SetText(1, v)
}

// MudpRequest_List is a list of MudpRequest.
type MudpRequest_List struct{ capnp.List }

// NewMudpRequest creates a new list of MudpRequest.
func NewMudpRequest_List(s *capnp.Segment, sz int32) (MudpRequest_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return MudpRequest_List{l}, err
}

func (s MudpRequest_List) At(i int) MudpRequest { return MudpRequest{s.List.Struct(i)} }

func (s MudpRequest_List) Set(i int, v MudpRequest) error { return s.List.SetStruct(i, v.Struct) }

func (s MudpRequest_List) String() string {
	str, _ := text.MarshalList(0xfd0cdac648e8286b, s.List)
	return str
}

// MudpRequest_Future is a wrapper for a MudpRequest promised by a client call.
type MudpRequest_Future struct{ *capnp.Future }

func (p MudpRequest_Future) Struct() (MudpRequest, error) {
	s, err := p.Future.Struct()
	return MudpRequest{s}, err
}

type MudpResponse struct{ capnp.Struct }

// MudpResponse_TypeID is the unique identifier for the type MudpResponse.
const MudpResponse_TypeID = 0x933e53369a6293b3

func NewMudpResponse(s *capnp.Segment) (MudpResponse, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return MudpResponse{st}, err
}

func NewRootMudpResponse(s *capnp.Segment) (MudpResponse, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return MudpResponse{st}, err
}

func ReadRootMudpResponse(msg *capnp.Message) (MudpResponse, error) {
	root, err := msg.Root()
	return MudpResponse{root.Struct()}, err
}

func (s MudpResponse) String() string {
	str, _ := text.Marshal(0x933e53369a6293b3, s.Struct)
	return str
}

func (s MudpResponse) Namespace() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s MudpResponse) HasNamespace() bool {
	return s.Struct.HasPtr(0)
}

func (s MudpResponse) NamespaceBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s MudpResponse) SetNamespace(v string) error {
	return s.Struct.SetText(0, v)
}

func (s MudpResponse) Envelope() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.DataList{List: p.List()}, err
}

func (s MudpResponse) HasEnvelope() bool {
	return s.Struct.HasPtr(1)
}

func (s MudpResponse) SetEnvelope(v capnp.DataList) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewEnvelope sets the envelope field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s MudpResponse) NewEnvelope(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// MudpResponse_List is a list of MudpResponse.
type MudpResponse_List struct{ capnp.List }

// NewMudpResponse creates a new list of MudpResponse.
func NewMudpResponse_List(s *capnp.Segment, sz int32) (MudpResponse_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return MudpResponse_List{l}, err
}

func (s MudpResponse_List) At(i int) MudpResponse { return MudpResponse{s.List.Struct(i)} }

func (s MudpResponse_List) Set(i int, v MudpResponse) error { return s.List.SetStruct(i, v.Struct) }

func (s MudpResponse_List) String() string {
	str, _ := text.MarshalList(0x933e53369a6293b3, s.List)
	return str
}

// MudpResponse_Future is a wrapper for a MudpResponse promised by a client call.
type MudpResponse_Future struct{ *capnp.Future }

func (p MudpResponse_Future) Struct() (MudpResponse, error) {
	s, err := p.Future.Struct()
	return MudpResponse{s}, err
}

type MudpPacket struct{ capnp.Struct }
type MudpPacket_Which uint16

const (
	MudpPacket_Which_request  MudpPacket_Which = 0
	MudpPacket_Which_response MudpPacket_Which = 1
)

func (w MudpPacket_Which) String() string {
	const s = "requestresponse"
	switch w {
	case MudpPacket_Which_request:
		return s[0:7]
	case MudpPacket_Which_response:
		return s[7:15]

	}
	return "MudpPacket_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// MudpPacket_TypeID is the unique identifier for the type MudpPacket.
const MudpPacket_TypeID = 0xf3610b67cd26c199

func NewMudpPacket(s *capnp.Segment) (MudpPacket, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return MudpPacket{st}, err
}

func NewRootMudpPacket(s *capnp.Segment) (MudpPacket, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return MudpPacket{st}, err
}

func ReadRootMudpPacket(msg *capnp.Message) (MudpPacket, error) {
	root, err := msg.Root()
	return MudpPacket{root.Struct()}, err
}

func (s MudpPacket) String() string {
	str, _ := text.Marshal(0xf3610b67cd26c199, s.Struct)
	return str
}

func (s MudpPacket) Which() MudpPacket_Which {
	return MudpPacket_Which(s.Struct.Uint16(0))
}
func (s MudpPacket) Request() (MudpRequest, error) {
	if s.Struct.Uint16(0) != 0 {
		panic("Which() != request")
	}
	p, err := s.Struct.Ptr(0)
	return MudpRequest{Struct: p.Struct()}, err
}

func (s MudpPacket) HasRequest() bool {
	if s.Struct.Uint16(0) != 0 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s MudpPacket) SetRequest(v MudpRequest) error {
	s.Struct.SetUint16(0, 0)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewRequest sets the request field to a newly
// allocated MudpRequest struct, preferring placement in s's segment.
func (s MudpPacket) NewRequest() (MudpRequest, error) {
	s.Struct.SetUint16(0, 0)
	ss, err := NewMudpRequest(s.Struct.Segment())
	if err != nil {
		return MudpRequest{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s MudpPacket) Response() (MudpResponse, error) {
	if s.Struct.Uint16(0) != 1 {
		panic("Which() != response")
	}
	p, err := s.Struct.Ptr(0)
	return MudpResponse{Struct: p.Struct()}, err
}

func (s MudpPacket) HasResponse() bool {
	if s.Struct.Uint16(0) != 1 {
		return false
	}
	return s.Struct.HasPtr(0)
}

func (s MudpPacket) SetResponse(v MudpResponse) error {
	s.Struct.SetUint16(0, 1)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewResponse sets the response field to a newly
// allocated MudpResponse struct, preferring placement in s's segment.
func (s MudpPacket) NewResponse() (MudpResponse, error) {
	s.Struct.SetUint16(0, 1)
	ss, err := NewMudpResponse(s.Struct.Segment())
	if err != nil {
		return MudpResponse{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// MudpPacket_List is a list of MudpPacket.
type MudpPacket_List struct{ capnp.List }

// NewMudpPacket creates a new list of MudpPacket.
func NewMudpPacket_List(s *capnp.Segment, sz int32) (MudpPacket_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return MudpPacket_List{l}, err
}

func (s MudpPacket_List) At(i int) MudpPacket { return MudpPacket{s.List.Struct(i)} }

func (s MudpPacket_List) Set(i int, v MudpPacket) error { return s.List.SetStruct(i, v.Struct) }

func (s MudpPacket_List) String() string {
	str, _ := text.MarshalList(0xf3610b67cd26c199, s.List)
	return str
}

// MudpPacket_Future is a wrapper for a MudpPacket promised by a client call.
type MudpPacket_Future struct{ *capnp.Future }

func (p MudpPacket_Future) Struct() (MudpPacket, error) {
	s, err := p.Future.Struct()
	return MudpPacket{s}, err
}

func (p MudpPacket_Future) Request() MudpRequest_Future {
	return MudpRequest_Future{Future: p.Future.Field(0, nil)}
}

func (p MudpPacket_Future) Response() MudpResponse_Future {
	return MudpResponse_Future{Future: p.Future.Field(0, nil)}
}

const schema_ef83879a531f9bf3 = "x\xdal\x90\xc1j\x1aQ\x18\x85\xcf\xb9\xa3\xb5\xe0\xd8" +
	"\x8eL\x97-Rh\x8b-\xb5(\x94R\xbahK " +
	"\x10B\x02\xde\xf1\x09&z\xc9\xc28\xde8\x1a\x97Y" +
	"d\x91\x07\x90,4\xe4%\x92]vq\x91]^ " +
	"O\x90]\x16\xae\xb207\xcch\x9c \xee\x0e\xf7\x1e" +
	"\xfe\xef\xff~g\xfc_T\xd2m\x01\xc8\xf7\xe9W\xe6" +
	"b\xb03\xfaU\xfb;@>K39-\xd4F\xc7" +
	"G\xf7H\x8b\x0c\xe0~\xe0\xd8\xfd\xcc(}d\x1f4" +
	"\xc3\xab/7\xbbY\x7f\x02\x99\xe5\x8b\xf2:3\x02p" +
	"Ox\xee\x9e\xc5\xeda\xdcn\x16\xef6\xaeo\xed\xe9" +
	"\xaa\xc9\x0f\xbct\x19\xa7)\xfb(\x99V\xaf\xa1\x7f\xd4" +
	"}-\x02\xfdg\xbb\xd7\xd0\x9e\x0au;\x08\x15\xaa\xa4" +
	"|m\xa5\x80\x14\x81\xfcW\x0f\x90E\x8b\xf2\xa7`\x9e" +
	"|\xc7\xe8\xb1\xb2\x09\xc8\xb2E\xb9%h\x02\xbf\xa5B" +
	"\xed\xd7AE\x1b\x826hTp\xa0\xf6\xdaZ\x01\xe0" +
	"\x1b\xb0j\x919\x88(.\xb8\x9cs\xab\x05\xbf\xdeT" +
	"\xdd9\xd56f\x86]\x03\xe4'\x8b\xb2,\x98\xe3\xa3" +
	"\x99qK\x11\xf7\xbbE\xf9[\xf0\xb0\xa3\xf6{*\xec" +
	"\xd2I\xb4A:\xa0\xe9<\xab\x00t\x92s\xcf\x7f\x97" +
	"\x17\xf0\xfe\xcd\x06-y\x7fK\x16Xx\x97\xbc\x84\xff" +
	"V+\xd5\x89\xadrXy\x84\xa7\x00\x00\x00\xff\xffI" +
	"*{z"

func init() {
	schemas.Register(schema_ef83879a531f9bf3,
		0x933e53369a6293b3,
		0xf3610b67cd26c199,
		0xfd0cdac648e8286b)
}
