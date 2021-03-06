// Code generated by protoc-gen-go.
// source: laptop_message.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import google_protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Laptop struct {
	Id          string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Brand       string                     `protobuf:"bytes,2,opt,name=brand" json:"brand,omitempty"`
	Name        string                     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Ram         *Memory                    `protobuf:"bytes,4,opt,name=ram" json:"ram,omitempty"`
	Cpu         *CPU                       `protobuf:"bytes,5,opt,name=cpu" json:"cpu,omitempty"`
	Gpus        []*GPU                     `protobuf:"bytes,6,rep,name=gpus" json:"gpus,omitempty"`
	Storages    []*Storage                 `protobuf:"bytes,7,rep,name=storages" json:"storages,omitempty"`
	Screen      *Screen                    `protobuf:"bytes,8,opt,name=screen" json:"screen,omitempty"`
	Keyboard    *Keyboard                  `protobuf:"bytes,9,opt,name=keyboard" json:"keyboard,omitempty"`
	WeightKg    float64                    `protobuf:"fixed64,10,opt,name=weight_kg" json:"weight_kg,omitempty"`
	WeightLb    float64                    `protobuf:"fixed64,11,opt,name=weight_lb" json:"weight_lb,omitempty"`
	PricingUsd  float64                    `protobuf:"fixed64,12,opt,name=pricing_usd" json:"pricing_usd,omitempty"`
	ReleaseYear uint32                     `protobuf:"varint,13,opt,name=release_year" json:"release_year,omitempty"`
	UpdatedAt   *google_protobuf.Timestamp `protobuf:"bytes,14,opt,name=updated_at" json:"updated_at,omitempty"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}

func (m *Laptop) GetRam() *Memory {
	if m != nil {
		return m.Ram
	}
	return nil
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

func (m *Laptop) GetUpdatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
}
