// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: order_management.proto

package ecommerce

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Items []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"` // repeated is used to represent the fields that can be repeated any number of times including zero in a
	// message.
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Destination string  `protobuf:"bytes,5,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_order_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_order_management_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Order) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Order) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type CombinedShipment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	OrderList []*Order `protobuf:"bytes,3,rep,name=orderList,proto3" json:"orderList,omitempty"`
}

func (x *CombinedShipment) Reset() {
	*x = CombinedShipment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CombinedShipment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinedShipment) ProtoMessage() {}

func (x *CombinedShipment) ProtoReflect() protoreflect.Message {
	mi := &file_order_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinedShipment.ProtoReflect.Descriptor instead.
func (*CombinedShipment) Descriptor() ([]byte, []int) {
	return file_order_management_proto_rawDescGZIP(), []int{1}
}

func (x *CombinedShipment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CombinedShipment) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CombinedShipment) GetOrderList() []*Order {
	if x != nil {
		return x.OrderList
	}
	return nil
}

var File_order_management_proto protoreflect.FileDescriptor

var file_order_management_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6a, 0x0a,
	0x10, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xdd, 0x02, 0x0a, 0x0f, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a,
	0x08, 0x61, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3a, 0x0a, 0x08, 0x67, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x1a, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x65, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x10, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x28, 0x01, 0x12, 0x4e, 0x0a, 0x0d, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1b, 0x2e, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x62, 0x69, 0x6e, 0x65, 0x64, 0x53, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x28, 0x01, 0x30, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_order_management_proto_rawDescOnce sync.Once
	file_order_management_proto_rawDescData = file_order_management_proto_rawDesc
)

func file_order_management_proto_rawDescGZIP() []byte {
	file_order_management_proto_rawDescOnce.Do(func() {
		file_order_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_management_proto_rawDescData)
	})
	return file_order_management_proto_rawDescData
}

var file_order_management_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_order_management_proto_goTypes = []interface{}{
	(*Order)(nil),                // 0: ecommerce.Order
	(*CombinedShipment)(nil),     // 1: ecommerce.CombinedShipment
	(*wrappers.StringValue)(nil), // 2: google.protobuf.StringValue
}
var file_order_management_proto_depIdxs = []int32{
	0, // 0: ecommerce.CombinedShipment.orderList:type_name -> ecommerce.Order
	0, // 1: ecommerce.OrderManagement.addOrder:input_type -> ecommerce.Order
	2, // 2: ecommerce.OrderManagement.getOrder:input_type -> google.protobuf.StringValue
	2, // 3: ecommerce.OrderManagement.searchOrders:input_type -> google.protobuf.StringValue
	0, // 4: ecommerce.OrderManagement.updateOrders:input_type -> ecommerce.Order
	2, // 5: ecommerce.OrderManagement.processOrders:input_type -> google.protobuf.StringValue
	2, // 6: ecommerce.OrderManagement.addOrder:output_type -> google.protobuf.StringValue
	0, // 7: ecommerce.OrderManagement.getOrder:output_type -> ecommerce.Order
	0, // 8: ecommerce.OrderManagement.searchOrders:output_type -> ecommerce.Order
	2, // 9: ecommerce.OrderManagement.updateOrders:output_type -> google.protobuf.StringValue
	1, // 10: ecommerce.OrderManagement.processOrders:output_type -> ecommerce.CombinedShipment
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_management_proto_init() }
func file_order_management_proto_init() {
	if File_order_management_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_order_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CombinedShipment); i {
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
			RawDescriptor: file_order_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_management_proto_goTypes,
		DependencyIndexes: file_order_management_proto_depIdxs,
		MessageInfos:      file_order_management_proto_msgTypes,
	}.Build()
	File_order_management_proto = out.File
	file_order_management_proto_rawDesc = nil
	file_order_management_proto_goTypes = nil
	file_order_management_proto_depIdxs = nil
}