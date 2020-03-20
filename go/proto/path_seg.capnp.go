// Code generated by capnpc-go. DO NOT EDIT.

package proto

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type PathSegment struct{ capnp.Struct }

// PathSegment_TypeID is the unique identifier for the type PathSegment.
const PathSegment_TypeID = 0x939cc09a86ba70fa

func NewPathSegment(s *capnp.Segment) (PathSegment, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PathSegment{st}, err
}

func NewRootPathSegment(s *capnp.Segment) (PathSegment, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return PathSegment{st}, err
}

func ReadRootPathSegment(msg *capnp.Message) (PathSegment, error) {
	root, err := msg.RootPtr()
	return PathSegment{root.Struct()}, err
}

func (s PathSegment) String() string {
	str, _ := text.Marshal(0x939cc09a86ba70fa, s.Struct)
	return str
}

func (s PathSegment) Sdata() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s PathSegment) HasSdata() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathSegment) SetSdata(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s PathSegment) AsEntries() (SignedBlob_List, error) {
	p, err := s.Struct.Ptr(1)
	return SignedBlob_List{List: p.List()}, err
}

func (s PathSegment) HasAsEntries() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s PathSegment) SetAsEntries(v SignedBlob_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewAsEntries sets the asEntries field to a newly
// allocated SignedBlob_List, preferring placement in s's segment.
func (s PathSegment) NewAsEntries(n int32) (SignedBlob_List, error) {
	l, err := NewSignedBlob_List(s.Struct.Segment(), n)
	if err != nil {
		return SignedBlob_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// PathSegment_List is a list of PathSegment.
type PathSegment_List struct{ capnp.List }

// NewPathSegment creates a new list of PathSegment.
func NewPathSegment_List(s *capnp.Segment, sz int32) (PathSegment_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return PathSegment_List{l}, err
}

func (s PathSegment_List) At(i int) PathSegment { return PathSegment{s.List.Struct(i)} }

func (s PathSegment_List) Set(i int, v PathSegment) error { return s.List.SetStruct(i, v.Struct) }

func (s PathSegment_List) String() string {
	str, _ := text.MarshalList(0x939cc09a86ba70fa, s.List)
	return str
}

// PathSegment_Promise is a wrapper for a PathSegment promised by a client call.
type PathSegment_Promise struct{ *capnp.Pipeline }

func (p PathSegment_Promise) Struct() (PathSegment, error) {
	s, err := p.Pipeline.Struct()
	return PathSegment{s}, err
}

type PathSegmentSignedData struct{ capnp.Struct }

// PathSegmentSignedData_TypeID is the unique identifier for the type PathSegmentSignedData.
const PathSegmentSignedData_TypeID = 0xc7cf7a18177aec2a

func NewPathSegmentSignedData(s *capnp.Segment) (PathSegmentSignedData, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PathSegmentSignedData{st}, err
}

func NewRootPathSegmentSignedData(s *capnp.Segment) (PathSegmentSignedData, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return PathSegmentSignedData{st}, err
}

func ReadRootPathSegmentSignedData(msg *capnp.Message) (PathSegmentSignedData, error) {
	root, err := msg.RootPtr()
	return PathSegmentSignedData{root.Struct()}, err
}

func (s PathSegmentSignedData) String() string {
	str, _ := text.Marshal(0xc7cf7a18177aec2a, s.Struct)
	return str
}

func (s PathSegmentSignedData) InfoF() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s PathSegmentSignedData) HasInfoF() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathSegmentSignedData) SetInfoF(v []byte) error {
	return s.Struct.SetData(0, v)
}

// PathSegmentSignedData_List is a list of PathSegmentSignedData.
type PathSegmentSignedData_List struct{ capnp.List }

// NewPathSegmentSignedData creates a new list of PathSegmentSignedData.
func NewPathSegmentSignedData_List(s *capnp.Segment, sz int32) (PathSegmentSignedData_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return PathSegmentSignedData_List{l}, err
}

func (s PathSegmentSignedData_List) At(i int) PathSegmentSignedData {
	return PathSegmentSignedData{s.List.Struct(i)}
}

func (s PathSegmentSignedData_List) Set(i int, v PathSegmentSignedData) error {
	return s.List.SetStruct(i, v.Struct)
}

func (s PathSegmentSignedData_List) String() string {
	str, _ := text.MarshalList(0xc7cf7a18177aec2a, s.List)
	return str
}

// PathSegmentSignedData_Promise is a wrapper for a PathSegmentSignedData promised by a client call.
type PathSegmentSignedData_Promise struct{ *capnp.Pipeline }

func (p PathSegmentSignedData_Promise) Struct() (PathSegmentSignedData, error) {
	s, err := p.Pipeline.Struct()
	return PathSegmentSignedData{s}, err
}

type ASEntry struct{ capnp.Struct }
type ASEntry_exts ASEntry

// ASEntry_TypeID is the unique identifier for the type ASEntry.
const ASEntry_TypeID = 0xd4a209e8e78874ff

func NewASEntry(s *capnp.Segment) (ASEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 4})
	return ASEntry{st}, err
}

func NewRootASEntry(s *capnp.Segment) (ASEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 32, PointerCount: 4})
	return ASEntry{st}, err
}

func ReadRootASEntry(msg *capnp.Message) (ASEntry, error) {
	root, err := msg.RootPtr()
	return ASEntry{root.Struct()}, err
}

func (s ASEntry) String() string {
	str, _ := text.Marshal(0xd4a209e8e78874ff, s.Struct)
	return str
}

func (s ASEntry) Isdas() uint64 {
	return s.Struct.Uint64(0)
}

func (s ASEntry) SetIsdas(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s ASEntry) TrcVer() uint64 {
	return s.Struct.Uint64(8)
}

func (s ASEntry) SetTrcVer(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s ASEntry) CertVer() uint64 {
	return s.Struct.Uint64(16)
}

func (s ASEntry) SetCertVer(v uint64) {
	s.Struct.SetUint64(16, v)
}

func (s ASEntry) IfIDSize() uint8 {
	return s.Struct.Uint8(24)
}

func (s ASEntry) SetIfIDSize(v uint8) {
	s.Struct.SetUint8(24, v)
}

func (s ASEntry) Hops() (HopEntry_List, error) {
	p, err := s.Struct.Ptr(0)
	return HopEntry_List{List: p.List()}, err
}

func (s ASEntry) HasHops() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ASEntry) SetHops(v HopEntry_List) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewHops sets the hops field to a newly
// allocated HopEntry_List, preferring placement in s's segment.
func (s ASEntry) NewHops(n int32) (HopEntry_List, error) {
	l, err := NewHopEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return HopEntry_List{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

func (s ASEntry) Mtu() uint16 {
	return s.Struct.Uint16(26)
}

func (s ASEntry) SetMtu(v uint16) {
	s.Struct.SetUint16(26, v)
}

func (s ASEntry) Exts() ASEntry_exts { return ASEntry_exts(s) }

func (s ASEntry_exts) RoutingPolicy() (RoutingPolicyExt, error) {
	p, err := s.Struct.Ptr(1)
	return RoutingPolicyExt{Struct: p.Struct()}, err
}

func (s ASEntry_exts) HasRoutingPolicy() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s ASEntry_exts) SetRoutingPolicy(v RoutingPolicyExt) error {
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewRoutingPolicy sets the routingPolicy field to a newly
// allocated RoutingPolicyExt struct, preferring placement in s's segment.
func (s ASEntry_exts) NewRoutingPolicy() (RoutingPolicyExt, error) {
	ss, err := NewRoutingPolicyExt(s.Struct.Segment())
	if err != nil {
		return RoutingPolicyExt{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s ASEntry_exts) Sibra() (SibraPCBExt, error) {
	p, err := s.Struct.Ptr(2)
	return SibraPCBExt{Struct: p.Struct()}, err
}

func (s ASEntry_exts) HasSibra() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s ASEntry_exts) SetSibra(v SibraPCBExt) error {
	return s.Struct.SetPtr(2, v.Struct.ToPtr())
}

// NewSibra sets the sibra field to a newly
// allocated SibraPCBExt struct, preferring placement in s's segment.
func (s ASEntry_exts) NewSibra() (SibraPCBExt, error) {
	ss, err := NewSibraPCBExt(s.Struct.Segment())
	if err != nil {
		return SibraPCBExt{}, err
	}
	err = s.Struct.SetPtr(2, ss.Struct.ToPtr())
	return ss, err
}

func (s ASEntry_exts) WatchDogMetric() (WatchDogMetricExt, error) {
	p, err := s.Struct.Ptr(3)
	return WatchDogMetricExt{Struct: p.Struct()}, err
}

func (s ASEntry_exts) HasWatchDogMetric() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s ASEntry_exts) SetWatchDogMetric(v WatchDogMetricExt) error {
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewWatchDogMetric sets the watchDogMetric field to a newly
// allocated WatchDogMetricExt struct, preferring placement in s's segment.
func (s ASEntry_exts) NewWatchDogMetric() (WatchDogMetricExt, error) {
	ss, err := NewWatchDogMetricExt(s.Struct.Segment())
	if err != nil {
		return WatchDogMetricExt{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// ASEntry_List is a list of ASEntry.
type ASEntry_List struct{ capnp.List }

// NewASEntry creates a new list of ASEntry.
func NewASEntry_List(s *capnp.Segment, sz int32) (ASEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 32, PointerCount: 4}, sz)
	return ASEntry_List{l}, err
}

func (s ASEntry_List) At(i int) ASEntry { return ASEntry{s.List.Struct(i)} }

func (s ASEntry_List) Set(i int, v ASEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s ASEntry_List) String() string {
	str, _ := text.MarshalList(0xd4a209e8e78874ff, s.List)
	return str
}

// ASEntry_Promise is a wrapper for a ASEntry promised by a client call.
type ASEntry_Promise struct{ *capnp.Pipeline }

func (p ASEntry_Promise) Struct() (ASEntry, error) {
	s, err := p.Pipeline.Struct()
	return ASEntry{s}, err
}

func (p ASEntry_Promise) Exts() ASEntry_exts_Promise { return ASEntry_exts_Promise{p.Pipeline} }

// ASEntry_exts_Promise is a wrapper for a ASEntry_exts promised by a client call.
type ASEntry_exts_Promise struct{ *capnp.Pipeline }

func (p ASEntry_exts_Promise) Struct() (ASEntry_exts, error) {
	s, err := p.Pipeline.Struct()
	return ASEntry_exts{s}, err
}

func (p ASEntry_exts_Promise) RoutingPolicy() RoutingPolicyExt_Promise {
	return RoutingPolicyExt_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p ASEntry_exts_Promise) Sibra() SibraPCBExt_Promise {
	return SibraPCBExt_Promise{Pipeline: p.Pipeline.GetPipeline(2)}
}

func (p ASEntry_exts_Promise) WatchDogMetric() WatchDogMetricExt_Promise {
	return WatchDogMetricExt_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

type HopEntry struct{ capnp.Struct }

// HopEntry_TypeID is the unique identifier for the type HopEntry.
const HopEntry_TypeID = 0x8bb1ddafb4872b0b

func NewHopEntry(s *capnp.Segment) (HopEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 1})
	return HopEntry{st}, err
}

func NewRootHopEntry(s *capnp.Segment) (HopEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 40, PointerCount: 1})
	return HopEntry{st}, err
}

func ReadRootHopEntry(msg *capnp.Message) (HopEntry, error) {
	root, err := msg.RootPtr()
	return HopEntry{root.Struct()}, err
}

func (s HopEntry) String() string {
	str, _ := text.Marshal(0x8bb1ddafb4872b0b, s.Struct)
	return str
}

func (s HopEntry) InIA() uint64 {
	return s.Struct.Uint64(0)
}

func (s HopEntry) SetInIA(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s HopEntry) RemoteInIF() uint64 {
	return s.Struct.Uint64(8)
}

func (s HopEntry) SetRemoteInIF(v uint64) {
	s.Struct.SetUint64(8, v)
}

func (s HopEntry) InMTU() uint16 {
	return s.Struct.Uint16(16)
}

func (s HopEntry) SetInMTU(v uint16) {
	s.Struct.SetUint16(16, v)
}

func (s HopEntry) OutIA() uint64 {
	return s.Struct.Uint64(24)
}

func (s HopEntry) SetOutIA(v uint64) {
	s.Struct.SetUint64(24, v)
}

func (s HopEntry) RemoteOutIF() uint64 {
	return s.Struct.Uint64(32)
}

func (s HopEntry) SetRemoteOutIF(v uint64) {
	s.Struct.SetUint64(32, v)
}

func (s HopEntry) HopF() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s HopEntry) HasHopF() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s HopEntry) SetHopF(v []byte) error {
	return s.Struct.SetData(0, v)
}

// HopEntry_List is a list of HopEntry.
type HopEntry_List struct{ capnp.List }

// NewHopEntry creates a new list of HopEntry.
func NewHopEntry_List(s *capnp.Segment, sz int32) (HopEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 40, PointerCount: 1}, sz)
	return HopEntry_List{l}, err
}

func (s HopEntry_List) At(i int) HopEntry { return HopEntry{s.List.Struct(i)} }

func (s HopEntry_List) Set(i int, v HopEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s HopEntry_List) String() string {
	str, _ := text.MarshalList(0x8bb1ddafb4872b0b, s.List)
	return str
}

// HopEntry_Promise is a wrapper for a HopEntry promised by a client call.
type HopEntry_Promise struct{ *capnp.Pipeline }

func (p HopEntry_Promise) Struct() (HopEntry, error) {
	s, err := p.Pipeline.Struct()
	return HopEntry{s}, err
}

type PCB struct{ capnp.Struct }

// PCB_TypeID is the unique identifier for the type PCB.
const PCB_TypeID = 0xd6c04763377951e5

func NewPCB(s *capnp.Segment) (PCB, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PCB{st}, err
}

func NewRootPCB(s *capnp.Segment) (PCB, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PCB{st}, err
}

func ReadRootPCB(msg *capnp.Message) (PCB, error) {
	root, err := msg.RootPtr()
	return PCB{root.Struct()}, err
}

func (s PCB) String() string {
	str, _ := text.Marshal(0xd6c04763377951e5, s.Struct)
	return str
}

func (s PCB) PathSeg() (PathSegment, error) {
	p, err := s.Struct.Ptr(0)
	return PathSegment{Struct: p.Struct()}, err
}

func (s PCB) HasPathSeg() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PCB) SetPathSeg(v PathSegment) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPathSeg sets the pathSeg field to a newly
// allocated PathSegment struct, preferring placement in s's segment.
func (s PCB) NewPathSeg() (PathSegment, error) {
	ss, err := NewPathSegment(s.Struct.Segment())
	if err != nil {
		return PathSegment{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PCB) IfID() uint64 {
	return s.Struct.Uint64(0)
}

func (s PCB) SetIfID(v uint64) {
	s.Struct.SetUint64(0, v)
}

// PCB_List is a list of PCB.
type PCB_List struct{ capnp.List }

// NewPCB creates a new list of PCB.
func NewPCB_List(s *capnp.Segment, sz int32) (PCB_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PCB_List{l}, err
}

func (s PCB_List) At(i int) PCB { return PCB{s.List.Struct(i)} }

func (s PCB_List) Set(i int, v PCB) error { return s.List.SetStruct(i, v.Struct) }

func (s PCB_List) String() string {
	str, _ := text.MarshalList(0xd6c04763377951e5, s.List)
	return str
}

// PCB_Promise is a wrapper for a PCB promised by a client call.
type PCB_Promise struct{ *capnp.Pipeline }

func (p PCB_Promise) Struct() (PCB, error) {
	s, err := p.Pipeline.Struct()
	return PCB{s}, err
}

func (p PCB_Promise) PathSeg() PathSegment_Promise {
	return PathSegment_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type PathSegMeta struct{ capnp.Struct }

// PathSegMeta_TypeID is the unique identifier for the type PathSegMeta.
const PathSegMeta_TypeID = 0x9f98567b3c0aba0f

func NewPathSegMeta(s *capnp.Segment) (PathSegMeta, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PathSegMeta{st}, err
}

func NewRootPathSegMeta(s *capnp.Segment) (PathSegMeta, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return PathSegMeta{st}, err
}

func ReadRootPathSegMeta(msg *capnp.Message) (PathSegMeta, error) {
	root, err := msg.RootPtr()
	return PathSegMeta{root.Struct()}, err
}

func (s PathSegMeta) String() string {
	str, _ := text.Marshal(0x9f98567b3c0aba0f, s.Struct)
	return str
}

func (s PathSegMeta) PathSeg() (PathSegment, error) {
	p, err := s.Struct.Ptr(0)
	return PathSegment{Struct: p.Struct()}, err
}

func (s PathSegMeta) HasPathSeg() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s PathSegMeta) SetPathSeg(v PathSegment) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewPathSeg sets the pathSeg field to a newly
// allocated PathSegment struct, preferring placement in s's segment.
func (s PathSegMeta) NewPathSeg() (PathSegment, error) {
	ss, err := NewPathSegment(s.Struct.Segment())
	if err != nil {
		return PathSegment{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s PathSegMeta) Type() PathSegType {
	return PathSegType(s.Struct.Uint16(0))
}

func (s PathSegMeta) SetType(v PathSegType) {
	s.Struct.SetUint16(0, uint16(v))
}

// PathSegMeta_List is a list of PathSegMeta.
type PathSegMeta_List struct{ capnp.List }

// NewPathSegMeta creates a new list of PathSegMeta.
func NewPathSegMeta_List(s *capnp.Segment, sz int32) (PathSegMeta_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return PathSegMeta_List{l}, err
}

func (s PathSegMeta_List) At(i int) PathSegMeta { return PathSegMeta{s.List.Struct(i)} }

func (s PathSegMeta_List) Set(i int, v PathSegMeta) error { return s.List.SetStruct(i, v.Struct) }

func (s PathSegMeta_List) String() string {
	str, _ := text.MarshalList(0x9f98567b3c0aba0f, s.List)
	return str
}

// PathSegMeta_Promise is a wrapper for a PathSegMeta promised by a client call.
type PathSegMeta_Promise struct{ *capnp.Pipeline }

func (p PathSegMeta_Promise) Struct() (PathSegMeta, error) {
	s, err := p.Pipeline.Struct()
	return PathSegMeta{s}, err
}

func (p PathSegMeta_Promise) PathSeg() PathSegment_Promise {
	return PathSegment_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type PathSegType uint16

// PathSegType_TypeID is the unique identifier for the type PathSegType.
const PathSegType_TypeID = 0xa1f7a9650aa23880

// Values of PathSegType.
const (
	PathSegType_unset PathSegType = 0
	PathSegType_up    PathSegType = 1
	PathSegType_down  PathSegType = 2
	PathSegType_core  PathSegType = 3
)

// String returns the enum's constant name.
func (c PathSegType) String() string {
	switch c {
	case PathSegType_unset:
		return "unset"
	case PathSegType_up:
		return "up"
	case PathSegType_down:
		return "down"
	case PathSegType_core:
		return "core"

	default:
		return ""
	}
}

// PathSegTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func PathSegTypeFromString(c string) PathSegType {
	switch c {
	case "unset":
		return PathSegType_unset
	case "up":
		return PathSegType_up
	case "down":
		return PathSegType_down
	case "core":
		return PathSegType_core

	default:
		return 0
	}
}

type PathSegType_List struct{ capnp.List }

func NewPathSegType_List(s *capnp.Segment, sz int32) (PathSegType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return PathSegType_List{l.List}, err
}

func (l PathSegType_List) At(i int) PathSegType {
	ul := capnp.UInt16List{List: l.List}
	return PathSegType(ul.At(i))
}

func (l PathSegType_List) Set(i int, v PathSegType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

const schema_fb8053d9fb34b837 = "x\xda\xacT]h\x1cU\x18=\xe7\xde\xd9\xdd4\x1a" +
	"g\x87]i*\xc2ZQ\xd0UKM+)EH" +
	"Z\xd36\x11\x82{\x9bX\x89(:\xddL7\x03f" +
	"v\xd9\xb9K\xba\x91\xd2\xfa\x10\x7f\x10\xf1A\x11--" +
	"i\x8b\x82\x81\x86Vi!\x0d\x0ai)\x14A\xf0A" +
	"\x11\x0d\x14\xd4V\xb4\xd0wiK\x1d\xb9\xd9\xec\x0fI" +
	"\xfb\xe6\xdb\x9d\x8f\xef\x9e\xef;\xe7\x9e3\x1bo\xb0W" +
	"<\x1d\xbb$\x00\xf5P,\x1e\xdd\xf3\xc4;gN]" +
	"\xfe\xea}(\x9b\xb1\xa8{n\xf3\xad_\x87\x0e\xddB" +
	"\x8c\x09 5\xc0\x0b)eN\x9b\x06\xf9\x12\xc1\xe8f" +
	"i\xfe\xed\xc3\x0bG>\x82c\xb3\xa5Y\x98\xe6\xe3b" +
	"15\xbbt\x9a\x11\x13`d\xcf\xb7?\xfb\xe6\x9eO" +
	"\xa7\x0d2W\"\xaf\x93\x8b\xa9G\xa59\xad\x97\xa6\xf9" +
	"\xd0\x96\x13\xed\xde\xcc?\xc7\xe1\xd8\xa2\xd9\x0b\xa6\x0e\xc8" +
	"\xc5\xd4{K\x8dS\xb2\x00F\xd1\xf4\xd4\xd1\x7f\xdb\xf5" +
	"\x05\xa8\xb5\xb4\xa2H\xbf\xfb\xd7\xdfkN\xfc\x84\xfb\xad" +
	"\x04\xcdd\xf9\x07\x98\x9a\x95\xa7\xc0({}rm\xe7" +
	"\xe4\x0f\x97V\xecZ\xa3ce\x99z\xd52\xb0#V" +
	"\x8f\x81\xad\x03)\x9bVK\xf7RK\xd5\x9aO\xbde" +
	"N\x9b\x0eX\x1f\x1a\x19\xfeT\xd5\xee\xfc\xae\x85\x9f\xef" +
	"H\xedv\xecp*\x167'\xc6\x0d\xb5\x92\xab\xc7^" +
	"\x0b\xbd\x82\xd8\x90wKAik\x7f\xb1\xb4#\xd0\xe5" +
	"*r\xa4\xea\x94\x16`\x11p>\xcb\x02\xeacIu" +
	"L\xd0!\xd34\xc5\xa3/\x03\xea\x88\xa4\xfaR\xd0\x11" +
	"mi\x0a\xc0\xf9\xa2\x0bP\xc7$\xd5IAG\xca4" +
	"%\xe0\xcc\x98\xe2\xe7\x92\xea\xb4\xa0cYiZ\x803" +
	"\xbb\x17P'%\xd5\x9c ci\xc6\x00\xe7\xac\x99s" +
	"ZR}#h\xfb\xc1\xc06\xae\x81\xe0\x1a0*{" +
	"\xe3E\xed\x0d\x04\x90\x03;\xeb\xc5\x8c\x1f\x0c\x0e\xbf\xc8" +
	"\x04\x04\x13`\xa6X\xd1\xab.\xbcPAB7o\xd8" +
	"c\xc5\xd2Nv@\xb0\xe3\x0e\xe4s\xae\x1e\x1b\xf2\x0a" +
	"\xe3\x9e\x0c\xb4\xe1\xdf\xd6\xe0\xff\xb8!\xf0\x88\xa4\xda\xd8" +
	"\xc2\xff\xa9\xdd\x80zRR\xf5\x0bf\xc2QW\xbb\x0d" +
	"d74*\xfa\x1e\x18\xf2>0'\xc9d\xb4\xfb\xca" +
	"\xcd\xee\xa9]]\xd3\x00M\xf1n\xe3\x07=\xa9\xdd\x15" +
	"\xe3\xb77\xc77\xa6\x1b\xa9\x1e\x93T\x9b\x05\x0f\x96j" +
	"W\x99l\x86\x00d\x12\xb4u\xb5\xe4\xd1nZ\x18\xa4" +
	"}\xf7\xd9\xc3UY\xf2\xcc\xec\xe4\xd2c\x1a\xda\xa4\xb3" +
	"\xfe\x01\x80\xc2Y\x97\x05(\x1d'\x0bd*A\xe8i" +
	"Y)\xd9\xa3\xc5\x89\xc0\xce\x17\xcb\xde*\xc8mCK" +
	"N\xda\xe0%\xf6\xebP\xdd+\xad$\xd3\x8c\x03\xce\x8e" +
	"2\xa0\xfa$U\xce\x88)\xd2\xc6\x98\xce\xa0Q\xb8_" +
	"R\x0d\x1b3\xc94\xdb\x00GM\x02*'\xa9^\x11" +
	"\x8c\xca\xc5\x8a\xf6\x83B\x0e\x99\xe2\x1b~\xbe\xcadt" +
	"\xe3\xea3[\xe6\x16\xcf\x7f\xb2\xcc5\x13\xfa{\xcb." +
	"\x93\xd1\x07\xaf_\x1b9w\xed\xf6\xdcr=\x9apu" +
	"~\xac\xafX@\xcf\xa0\xa7\xcb~\x9e\xc9h\xc4\xfa\xfa" +
	"\xc7\xde+\xbd\xd7\xeb-\xf5\xe5\xe5J+\x04z\xc8/" +
	"\x04\xde\xa8\xdd\xe7\xd6^\xc5j\xbcJ\x87Y\xb9MR" +
	"\xa5\x85\xb1\xe2\xbe\xe2jg\xb1\xaeEOM\x0c\x03\xf0" +
	"`\x03\xe0lW\xd3\xed\x0dW\x9d\xdb\x0a\xa83\x92j" +
	"\xc1\x08!j\xa9\xfa\xd6\x18`NR]4\xa9\xea\xac" +
	"\xa5\xea\xfc\xf3\x80Z\x90T\xdf\x0br9T\xdf\x19W" +
	"\\\x94T\xbf\x09:\xb1\x8eZ\xaa.?\x0c\xa8_$" +
	"\xd5UA\xc6\xd9\xf2\x9fr~\xcfBd\xfcp\xd4\x0d" +
	"\xeb\x11\xe9\xd1\xe5\xfc\x1e\xaf\\\xff<\x98\xf7\xca\xba\xe5" +
	";\xf2\xf7\x0d\xf4\x0d\xf9\x93\x1e\x00\xc6!\x18\xaf\xa5\xaa" +
	"\xc5\xe7\x8d\xbfv\xcd\xe7\x89q]\xa9g\xd4\xf6\xf6\xeb" +
	"p\x95:9\xf9\xdc\xf6\xff\xcd\xf0f\xbf\xfa\xb2\xff\x05" +
	"\x00\x00\xff\xff\x0e*\x8d\xa1"

func init() {
	schemas.Register(schema_fb8053d9fb34b837,
		0x8bb1ddafb4872b0b,
		0x939cc09a86ba70fa,
		0x9f98567b3c0aba0f,
		0xa1f7a9650aa23880,
		0xc2740afe9d859fff,
		0xc7cf7a18177aec2a,
		0xd4a209e8e78874ff,
		0xd6c04763377951e5)
}
