// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: protos/accounting-service.proto

package protos

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

type CreateLedgerEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InvoiceId   string  `protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	TypeId      string  `protobuf:"bytes,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Amount      float64 `protobuf:"fixed64,4,opt,name=amount,proto3" json:"amount,omitempty"`
	Date        string  `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *CreateLedgerEntryRequest) Reset() {
	*x = CreateLedgerEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLedgerEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLedgerEntryRequest) ProtoMessage() {}

func (x *CreateLedgerEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLedgerEntryRequest.ProtoReflect.Descriptor instead.
func (*CreateLedgerEntryRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateLedgerEntryRequest) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *CreateLedgerEntryRequest) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *CreateLedgerEntryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateLedgerEntryRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateLedgerEntryRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type CreateLedgerEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateLedgerEntryResponse) Reset() {
	*x = CreateLedgerEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLedgerEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLedgerEntryResponse) ProtoMessage() {}

func (x *CreateLedgerEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLedgerEntryResponse.ProtoReflect.Descriptor instead.
func (*CreateLedgerEntryResponse) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLedgerEntryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLedgerEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLedgerEntryRequest) Reset() {
	*x = GetLedgerEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedgerEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedgerEntryRequest) ProtoMessage() {}

func (x *GetLedgerEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedgerEntryRequest.ProtoReflect.Descriptor instead.
func (*GetLedgerEntryRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLedgerEntryRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLedgerEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ledger *Ledger `protobuf:"bytes,1,opt,name=ledger,proto3" json:"ledger,omitempty"`
}

func (x *GetLedgerEntryResponse) Reset() {
	*x = GetLedgerEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLedgerEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLedgerEntryResponse) ProtoMessage() {}

func (x *GetLedgerEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLedgerEntryResponse.ProtoReflect.Descriptor instead.
func (*GetLedgerEntryResponse) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLedgerEntryResponse) GetLedger() *Ledger {
	if x != nil {
		return x.Ledger
	}
	return nil
}

type UpdateLedgerEntryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ledger *Ledger `protobuf:"bytes,1,opt,name=ledger,proto3" json:"ledger,omitempty"`
}

func (x *UpdateLedgerEntryRequest) Reset() {
	*x = UpdateLedgerEntryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLedgerEntryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLedgerEntryRequest) ProtoMessage() {}

func (x *UpdateLedgerEntryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLedgerEntryRequest.ProtoReflect.Descriptor instead.
func (*UpdateLedgerEntryRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateLedgerEntryRequest) GetLedger() *Ledger {
	if x != nil {
		return x.Ledger
	}
	return nil
}

type UpdateLedgerEntryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ledger *Ledger `protobuf:"bytes,1,opt,name=ledger,proto3" json:"ledger,omitempty"`
}

func (x *UpdateLedgerEntryResponse) Reset() {
	*x = UpdateLedgerEntryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateLedgerEntryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateLedgerEntryResponse) ProtoMessage() {}

func (x *UpdateLedgerEntryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateLedgerEntryResponse.ProtoReflect.Descriptor instead.
func (*UpdateLedgerEntryResponse) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateLedgerEntryResponse) GetLedger() *Ledger {
	if x != nil {
		return x.Ledger
	}
	return nil
}

type ListLedgerEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *ListLedgerEntriesRequest) Reset() {
	*x = ListLedgerEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLedgerEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLedgerEntriesRequest) ProtoMessage() {}

func (x *ListLedgerEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLedgerEntriesRequest.ProtoReflect.Descriptor instead.
func (*ListLedgerEntriesRequest) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListLedgerEntriesRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *ListLedgerEntriesRequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type ListLedgerEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Ledger `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ListLedgerEntriesResponse) Reset() {
	*x = ListLedgerEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLedgerEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLedgerEntriesResponse) ProtoMessage() {}

func (x *ListLedgerEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLedgerEntriesResponse.ProtoReflect.Descriptor instead.
func (*ListLedgerEntriesResponse) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListLedgerEntriesResponse) GetEntries() []*Ledger {
	if x != nil {
		return x.Entries
	}
	return nil
}

type Ledger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InvoiceId   string  `protobuf:"bytes,2,opt,name=invoice_id,json=invoiceId,proto3" json:"invoice_id,omitempty"`
	TypeId      string  `protobuf:"bytes,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Amount      float64 `protobuf:"fixed64,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Date        string  `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *Ledger) Reset() {
	*x = Ledger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_accounting_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ledger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ledger) ProtoMessage() {}

func (x *Ledger) ProtoReflect() protoreflect.Message {
	mi := &file_protos_accounting_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ledger.ProtoReflect.Descriptor instead.
func (*Ledger) Descriptor() ([]byte, []int) {
	return file_protos_accounting_service_proto_rawDescGZIP(), []int{8}
}

func (x *Ledger) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ledger) GetInvoiceId() string {
	if x != nil {
		return x.InvoiceId
	}
	return ""
}

func (x *Ledger) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *Ledger) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Ledger) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Ledger) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_protos_accounting_service_proto protoreflect.FileDescriptor

var file_protos_accounting_service_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xa0, 0x01,
	0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x22, 0x2b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x22, 0x47, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x06, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x06, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x22, 0x54, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x22, 0x49, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x9e,
	0x01, 0x0a, 0x06, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x32,
	0x92, 0x03, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x60, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_accounting_service_proto_rawDescOnce sync.Once
	file_protos_accounting_service_proto_rawDescData = file_protos_accounting_service_proto_rawDesc
)

func file_protos_accounting_service_proto_rawDescGZIP() []byte {
	file_protos_accounting_service_proto_rawDescOnce.Do(func() {
		file_protos_accounting_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_accounting_service_proto_rawDescData)
	})
	return file_protos_accounting_service_proto_rawDescData
}

var file_protos_accounting_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_accounting_service_proto_goTypes = []any{
	(*CreateLedgerEntryRequest)(nil),  // 0: accounting.CreateLedgerEntryRequest
	(*CreateLedgerEntryResponse)(nil), // 1: accounting.CreateLedgerEntryResponse
	(*GetLedgerEntryRequest)(nil),     // 2: accounting.GetLedgerEntryRequest
	(*GetLedgerEntryResponse)(nil),    // 3: accounting.GetLedgerEntryResponse
	(*UpdateLedgerEntryRequest)(nil),  // 4: accounting.UpdateLedgerEntryRequest
	(*UpdateLedgerEntryResponse)(nil), // 5: accounting.UpdateLedgerEntryResponse
	(*ListLedgerEntriesRequest)(nil),  // 6: accounting.ListLedgerEntriesRequest
	(*ListLedgerEntriesResponse)(nil), // 7: accounting.ListLedgerEntriesResponse
	(*Ledger)(nil),                    // 8: accounting.Ledger
}
var file_protos_accounting_service_proto_depIdxs = []int32{
	8, // 0: accounting.GetLedgerEntryResponse.ledger:type_name -> accounting.Ledger
	8, // 1: accounting.UpdateLedgerEntryRequest.ledger:type_name -> accounting.Ledger
	8, // 2: accounting.UpdateLedgerEntryResponse.ledger:type_name -> accounting.Ledger
	8, // 3: accounting.ListLedgerEntriesResponse.entries:type_name -> accounting.Ledger
	0, // 4: accounting.AccountingService.CreateLedgerEntry:input_type -> accounting.CreateLedgerEntryRequest
	2, // 5: accounting.AccountingService.GetLedgerEntry:input_type -> accounting.GetLedgerEntryRequest
	4, // 6: accounting.AccountingService.UpdateLedgerEntry:input_type -> accounting.UpdateLedgerEntryRequest
	6, // 7: accounting.AccountingService.ListLedgerEntries:input_type -> accounting.ListLedgerEntriesRequest
	1, // 8: accounting.AccountingService.CreateLedgerEntry:output_type -> accounting.CreateLedgerEntryResponse
	3, // 9: accounting.AccountingService.GetLedgerEntry:output_type -> accounting.GetLedgerEntryResponse
	5, // 10: accounting.AccountingService.UpdateLedgerEntry:output_type -> accounting.UpdateLedgerEntryResponse
	7, // 11: accounting.AccountingService.ListLedgerEntries:output_type -> accounting.ListLedgerEntriesResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_accounting_service_proto_init() }
func file_protos_accounting_service_proto_init() {
	if File_protos_accounting_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_accounting_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLedgerEntryRequest); i {
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
		file_protos_accounting_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLedgerEntryResponse); i {
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
		file_protos_accounting_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetLedgerEntryRequest); i {
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
		file_protos_accounting_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetLedgerEntryResponse); i {
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
		file_protos_accounting_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateLedgerEntryRequest); i {
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
		file_protos_accounting_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateLedgerEntryResponse); i {
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
		file_protos_accounting_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListLedgerEntriesRequest); i {
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
		file_protos_accounting_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListLedgerEntriesResponse); i {
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
		file_protos_accounting_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Ledger); i {
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
			RawDescriptor: file_protos_accounting_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_accounting_service_proto_goTypes,
		DependencyIndexes: file_protos_accounting_service_proto_depIdxs,
		MessageInfos:      file_protos_accounting_service_proto_msgTypes,
	}.Build()
	File_protos_accounting_service_proto = out.File
	file_protos_accounting_service_proto_rawDesc = nil
	file_protos_accounting_service_proto_goTypes = nil
	file_protos_accounting_service_proto_depIdxs = nil
}
