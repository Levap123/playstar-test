// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/weather.proto

package proto

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

type GetWeatherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float32 `protobuf:"fixed32,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float32 `protobuf:"fixed32,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *GetWeatherRequest) Reset() {
	*x = GetWeatherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherRequest) ProtoMessage() {}

func (x *GetWeatherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherRequest.ProtoReflect.Descriptor instead.
func (*GetWeatherRequest) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{0}
}

func (x *GetWeatherRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GetWeatherRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type GetWeatherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temperature              float32 `protobuf:"fixed32,1,opt,name=temperature,proto3" json:"temperature,omitempty"`
	TemperatureApparent      float32 `protobuf:"fixed32,2,opt,name=temperature_apparent,json=temperatureApparent,proto3" json:"temperature_apparent,omitempty"`
	Humidity                 uint64  `protobuf:"varint,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	PrecipitationProbability uint64  `protobuf:"varint,4,opt,name=precipitation_probability,json=precipitationProbability,proto3" json:"precipitation_probability,omitempty"`
}

func (x *GetWeatherResponse) Reset() {
	*x = GetWeatherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_weather_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWeatherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWeatherResponse) ProtoMessage() {}

func (x *GetWeatherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_weather_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWeatherResponse.ProtoReflect.Descriptor instead.
func (*GetWeatherResponse) Descriptor() ([]byte, []int) {
	return file_proto_weather_proto_rawDescGZIP(), []int{1}
}

func (x *GetWeatherResponse) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *GetWeatherResponse) GetTemperatureApparent() float32 {
	if x != nil {
		return x.TemperatureApparent
	}
	return 0
}

func (x *GetWeatherResponse) GetHumidity() uint64 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *GetWeatherResponse) GetPrecipitationProbability() uint64 {
	if x != nil {
		return x.PrecipitationProbability
	}
	return 0
}

var File_proto_weather_proto protoreflect.FileDescriptor

var file_proto_weather_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x31, 0x0a, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x41,
	0x70, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x19, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x18, 0x70, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x32, 0x53, 0x0a, 0x0e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_weather_proto_rawDescOnce sync.Once
	file_proto_weather_proto_rawDescData = file_proto_weather_proto_rawDesc
)

func file_proto_weather_proto_rawDescGZIP() []byte {
	file_proto_weather_proto_rawDescOnce.Do(func() {
		file_proto_weather_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_weather_proto_rawDescData)
	})
	return file_proto_weather_proto_rawDescData
}

var file_proto_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_weather_proto_goTypes = []interface{}{
	(*GetWeatherRequest)(nil),  // 0: proto.GetWeatherRequest
	(*GetWeatherResponse)(nil), // 1: proto.GetWeatherResponse
}
var file_proto_weather_proto_depIdxs = []int32{
	0, // 0: proto.WeatherService.GetWeather:input_type -> proto.GetWeatherRequest
	1, // 1: proto.WeatherService.GetWeather:output_type -> proto.GetWeatherResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_weather_proto_init() }
func file_proto_weather_proto_init() {
	if File_proto_weather_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_weather_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherRequest); i {
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
		file_proto_weather_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWeatherResponse); i {
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
			RawDescriptor: file_proto_weather_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_weather_proto_goTypes,
		DependencyIndexes: file_proto_weather_proto_depIdxs,
		MessageInfos:      file_proto_weather_proto_msgTypes,
	}.Build()
	File_proto_weather_proto = out.File
	file_proto_weather_proto_rawDesc = nil
	file_proto_weather_proto_goTypes = nil
	file_proto_weather_proto_depIdxs = nil
}