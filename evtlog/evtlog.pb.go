// ----------------------------------------------------------------------------
// evtlog.proto - port data protobuf definitions
//
// January 2021
//
// Copyright (c) 2019, 2021 by Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ----------------------------------------------------------------------------

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: evtlog.proto

package evtlog

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Evtlog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version       string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	FileType      string `protobuf:"bytes,2,opt,name=file_type,json=fileType,proto3" json:"file_type,omitempty"`
	Compressed    bool   `protobuf:"varint,3,opt,name=compressed,proto3" json:"compressed,omitempty"`
	SeqNo         uint64 `protobuf:"varint,4,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	ChunckTotal   uint32 `protobuf:"varint,5,opt,name=chunck_total,json=chunckTotal,proto3" json:"chunck_total,omitempty"`
	ChunckNo      uint32 `protobuf:"varint,6,opt,name=chunck_no,json=chunckNo,proto3" json:"chunck_no,omitempty"`
	FileName      string `protobuf:"bytes,7,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize      uint64 `protobuf:"varint,8,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	FileMd5Chksum []byte `protobuf:"bytes,9,opt,name=file_md5_chksum,json=fileMd5Chksum,proto3" json:"file_md5_chksum,omitempty"`
	File          []byte `protobuf:"bytes,10,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *Evtlog) Reset() {
	*x = Evtlog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_evtlog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Evtlog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Evtlog) ProtoMessage() {}

func (x *Evtlog) ProtoReflect() protoreflect.Message {
	mi := &file_evtlog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Evtlog.ProtoReflect.Descriptor instead.
func (*Evtlog) Descriptor() ([]byte, []int) {
	return file_evtlog_proto_rawDescGZIP(), []int{0}
}

func (x *Evtlog) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Evtlog) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *Evtlog) GetCompressed() bool {
	if x != nil {
		return x.Compressed
	}
	return false
}

func (x *Evtlog) GetSeqNo() uint64 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *Evtlog) GetChunckTotal() uint32 {
	if x != nil {
		return x.ChunckTotal
	}
	return 0
}

func (x *Evtlog) GetChunckNo() uint32 {
	if x != nil {
		return x.ChunckNo
	}
	return 0
}

func (x *Evtlog) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Evtlog) GetFileSize() uint64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Evtlog) GetFileMd5Chksum() []byte {
	if x != nil {
		return x.FileMd5Chksum
	}
	return nil
}

func (x *Evtlog) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

var File_evtlog_proto protoreflect.FileDescriptor

var file_evtlog_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x74, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac,
	0x02, 0x0a, 0x06, 0x45, 0x76, 0x74, 0x6c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x73, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x75, 0x6e, 0x63,
	0x6b, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63,
	0x68, 0x75, 0x6e, 0x63, 0x6b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68,
	0x75, 0x6e, 0x63, 0x6b, 0x5f, 0x6e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63,
	0x68, 0x75, 0x6e, 0x63, 0x6b, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x64, 0x35, 0x5f, 0x63, 0x68,
	0x6b, 0x73, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65,
	0x4d, 0x64, 0x35, 0x43, 0x68, 0x6b, 0x73, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x0d, 0x5a,
	0x08, 0x2e, 0x2f, 0x65, 0x76, 0x74, 0x6c, 0x6f, 0x67, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_evtlog_proto_rawDescOnce sync.Once
	file_evtlog_proto_rawDescData = file_evtlog_proto_rawDesc
)

func file_evtlog_proto_rawDescGZIP() []byte {
	file_evtlog_proto_rawDescOnce.Do(func() {
		file_evtlog_proto_rawDescData = protoimpl.X.CompressGZIP(file_evtlog_proto_rawDescData)
	})
	return file_evtlog_proto_rawDescData
}

var file_evtlog_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_evtlog_proto_goTypes = []interface{}{
	(*Evtlog)(nil), // 0: Evtlog
}
var file_evtlog_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_evtlog_proto_init() }
func file_evtlog_proto_init() {
	if File_evtlog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_evtlog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Evtlog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_evtlog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_evtlog_proto_goTypes,
		DependencyIndexes: file_evtlog_proto_depIdxs,
		MessageInfos:      file_evtlog_proto_msgTypes,
	}.Build()
	File_evtlog_proto = out.File
	file_evtlog_proto_rawDesc = nil
	file_evtlog_proto_goTypes = nil
	file_evtlog_proto_depIdxs = nil
}
