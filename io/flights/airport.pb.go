//*
// Airport Records
//
// Airport records allow a carrier to control how airport names are displayed on the pass, and to specify GPS coordinates that are more appropriate to the carrier's location(s).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: io/flights/airport.proto

package flights

import (
	reflect "reflect"
	sync "sync"

	io "github.com/PassKit/passkit-golang-grpc-sdk/io"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Port is an optional record that allows the carrier to overwrite default airport names and their localizations.  A port can represent an origin, destination or transit port of a direct flight with stops.  If a port record does not exist, this information will be automatically populated with publicly available data.
type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IATA code of the port. At least one of IATA or ICAO airport code is required.
	// @tag: validateGeneric:"required_without=IcaoAirportCode,fixedLenAlphaNum=3"
	IataAirportCode string `protobuf:"bytes,1,opt,name=iataAirportCode,proto3" json:"iataAirportCode,omitempty" validateGeneric:"required_without=IcaoAirportCode,fixedLenAlphaNum=3"`
	// The IATA code of the port. At least one of IATA or ICAO airport code is required.
	// @tag: validateGeneric:"required_without=IataAirportCode,fixedLenAlphaNum=4|isdefault"
	IcaoAirportCode string `protobuf:"bytes,2,opt,name=icaoAirportCode,proto3" json:"icaoAirportCode,omitempty" validateGeneric:"required_without=IataAirportCode,fixedLenAlphaNum=4|isdefault"`
	// The name of the city associated with the airport can be used in back/text fields.
	// @tag: validateGeneric:"required"
	CityName string `protobuf:"bytes,3,opt,name=cityName,proto3" json:"cityName,omitempty" validateGeneric:"required"`
	// The localized name of the city to be displayed on the boarding pass.
	LocalizedCityName *io.LocalizedString `protobuf:"bytes,4,opt,name=localizedCityName,proto3" json:"localizedCityName,omitempty"`
	// The name of the airport to be displayed on the boarding pass above the airport code.
	// @tag: validateGeneric:"required"
	AirportName string `protobuf:"bytes,5,opt,name=airportName,proto3" json:"airportName,omitempty" validateGeneric:"required"`
	// The localized name of the airport to be displayed on the boarding pass above the airport code.
	LocalizedAirportName *io.LocalizedString `protobuf:"bytes,6,opt,name=localizedAirportName,proto3" json:"localizedAirportName,omitempty"`
	// The ISO 3166 country code of the port.
	// @tag: validateGeneric:"required,alpha,len=2"
	CountryCode string `protobuf:"bytes,7,opt,name=countryCode,proto3" json:"countryCode,omitempty" validateGeneric:"required,alpha,len=2"`
	// The timezone of the airport in IANA timezone format. This is required to ensure the correct rendering of times and dates in the time local to the port.
	// @tag: validateGeneric:"required,ianaTimeZone"
	Timezone string `protobuf:"bytes,8,opt,name=timezone,proto3" json:"timezone,omitempty" validateGeneric:"required,ianaTimeZone"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_flights_airport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_io_flights_airport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_io_flights_airport_proto_rawDescGZIP(), []int{0}
}

func (x *Port) GetIataAirportCode() string {
	if x != nil {
		return x.IataAirportCode
	}
	return ""
}

func (x *Port) GetIcaoAirportCode() string {
	if x != nil {
		return x.IcaoAirportCode
	}
	return ""
}

func (x *Port) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Port) GetLocalizedCityName() *io.LocalizedString {
	if x != nil {
		return x.LocalizedCityName
	}
	return nil
}

func (x *Port) GetAirportName() string {
	if x != nil {
		return x.AirportName
	}
	return ""
}

func (x *Port) GetLocalizedAirportName() *io.LocalizedString {
	if x != nil {
		return x.LocalizedAirportName
	}
	return nil
}

func (x *Port) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Port) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

// Airport Request is used for retrieving or deleting a port record.
type AirportCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @tag: validate:"required"
	// The IATA or ICAO airport code.
	AirportCode string `protobuf:"bytes,1,opt,name=airportCode,proto3" json:"airportCode,omitempty" validate:"required"`
}

func (x *AirportCode) Reset() {
	*x = AirportCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_io_flights_airport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AirportCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AirportCode) ProtoMessage() {}

func (x *AirportCode) ProtoReflect() protoreflect.Message {
	mi := &file_io_flights_airport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AirportCode.ProtoReflect.Descriptor instead.
func (*AirportCode) Descriptor() ([]byte, []int) {
	return file_io_flights_airport_proto_rawDescGZIP(), []int{1}
}

func (x *AirportCode) GetAirportCode() string {
	if x != nil {
		return x.AirportCode
	}
	return ""
}

var File_io_flights_airport_proto protoreflect.FileDescriptor

var file_io_flights_airport_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x2f, 0x61, 0x69, 0x72,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x66, 0x6c, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x1a, 0x1c, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa1, 0x04, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x2f, 0x0a, 0x0f, 0x69, 0x61,
	0x74, 0x61, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x05, 0x92, 0x41, 0x02, 0x78, 0x03, 0x52, 0x0f, 0x69, 0x61, 0x74, 0x61,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x0f, 0x69,
	0x63, 0x61, 0x6f, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x05, 0x92, 0x41, 0x02, 0x78, 0x04, 0x52, 0x0f, 0x69, 0x63, 0x61,
	0x6f, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x11, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x11, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x64, 0x43, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a,
	0x14, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x69, 0x6f,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x14, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x41, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65,
	0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0xae, 0x01, 0x92, 0x41, 0xaa, 0x01, 0x0a, 0xa7, 0x01, 0x2a, 0x0e,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x32, 0x51,
	0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x20,
	0x61, 0x72, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x65, 0x61, 0x63, 0x68, 0x20, 0x70, 0x6f, 0x72, 0x74, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20,
	0x61, 0x20, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x20, 0x69, 0x6e, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6f, 0x66,
	0x2e, 0xd2, 0x01, 0x0f, 0x69, 0x61, 0x74, 0x61, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0xd2, 0x01, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0xd2, 0x01,
	0x0b, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x0b, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0xd2, 0x01, 0x08, 0x74, 0x69, 0x6d,
	0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x2f, 0x0a, 0x0b, 0x41, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x69, 0x72, 0x70, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x69, 0x72, 0x70, 0x6f,
	0x72, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x5f, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x6b, 0x69, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6c, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x5a, 0x2c, 0x73, 0x74, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x6b, 0x69,
	0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6f, 0x2f, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73,
	0xaa, 0x02, 0x14, 0x50, 0x61, 0x73, 0x73, 0x4b, 0x69, 0x74, 0x2e, 0x47, 0x72, 0x70, 0x63, 0x2e,
	0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_io_flights_airport_proto_rawDescOnce sync.Once
	file_io_flights_airport_proto_rawDescData = file_io_flights_airport_proto_rawDesc
)

func file_io_flights_airport_proto_rawDescGZIP() []byte {
	file_io_flights_airport_proto_rawDescOnce.Do(func() {
		file_io_flights_airport_proto_rawDescData = protoimpl.X.CompressGZIP(file_io_flights_airport_proto_rawDescData)
	})
	return file_io_flights_airport_proto_rawDescData
}

var file_io_flights_airport_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_io_flights_airport_proto_goTypes = []interface{}{
	(*Port)(nil),               // 0: flights.Port
	(*AirportCode)(nil),        // 1: flights.AirportCode
	(*io.LocalizedString)(nil), // 2: io.LocalizedString
}
var file_io_flights_airport_proto_depIdxs = []int32{
	2, // 0: flights.Port.localizedCityName:type_name -> io.LocalizedString
	2, // 1: flights.Port.localizedAirportName:type_name -> io.LocalizedString
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_io_flights_airport_proto_init() }
func file_io_flights_airport_proto_init() {
	if File_io_flights_airport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_io_flights_airport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
		file_io_flights_airport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AirportCode); i {
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
			RawDescriptor: file_io_flights_airport_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_io_flights_airport_proto_goTypes,
		DependencyIndexes: file_io_flights_airport_proto_depIdxs,
		MessageInfos:      file_io_flights_airport_proto_msgTypes,
	}.Build()
	File_io_flights_airport_proto = out.File
	file_io_flights_airport_proto_rawDesc = nil
	file_io_flights_airport_proto_goTypes = nil
	file_io_flights_airport_proto_depIdxs = nil
}
