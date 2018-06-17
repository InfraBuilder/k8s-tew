// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/firestore/admin/v1beta1/firestore_admin.proto

package admin // import "google.golang.org/genproto/googleapis/firestore/admin/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import longrunning "google.golang.org/genproto/googleapis/longrunning"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The type of index operation.
type IndexOperationMetadata_OperationType int32

const (
	// Unspecified. Never set by server.
	IndexOperationMetadata_OPERATION_TYPE_UNSPECIFIED IndexOperationMetadata_OperationType = 0
	// The operation is creating the index. Initiated by a `CreateIndex` call.
	IndexOperationMetadata_CREATING_INDEX IndexOperationMetadata_OperationType = 1
)

var IndexOperationMetadata_OperationType_name = map[int32]string{
	0: "OPERATION_TYPE_UNSPECIFIED",
	1: "CREATING_INDEX",
}
var IndexOperationMetadata_OperationType_value = map[string]int32{
	"OPERATION_TYPE_UNSPECIFIED": 0,
	"CREATING_INDEX":             1,
}

func (x IndexOperationMetadata_OperationType) String() string {
	return proto.EnumName(IndexOperationMetadata_OperationType_name, int32(x))
}
func (IndexOperationMetadata_OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{0, 0}
}

// Metadata for index operations. This metadata populates
// the metadata field of [google.longrunning.Operation][google.longrunning.Operation].
type IndexOperationMetadata struct {
	// The time that work began on the operation.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// The time the operation ended, either successfully or otherwise. Unset if
	// the operation is still active.
	EndTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	// The index resource that this operation is acting on. For example:
	// `projects/{project_id}/databases/{database_id}/indexes/{index_id}`
	Index string `protobuf:"bytes,3,opt,name=index" json:"index,omitempty"`
	// The type of index operation.
	OperationType IndexOperationMetadata_OperationType `protobuf:"varint,4,opt,name=operation_type,json=operationType,enum=google.firestore.admin.v1beta1.IndexOperationMetadata_OperationType" json:"operation_type,omitempty"`
	// True if the [google.longrunning.Operation] was cancelled. If the
	// cancellation is in progress, cancelled will be true but
	// [google.longrunning.Operation.done][google.longrunning.Operation.done] will be false.
	Cancelled bool `protobuf:"varint,5,opt,name=cancelled" json:"cancelled,omitempty"`
	// Progress of the existing operation, measured in number of documents.
	DocumentProgress     *Progress `protobuf:"bytes,6,opt,name=document_progress,json=documentProgress" json:"document_progress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IndexOperationMetadata) Reset()         { *m = IndexOperationMetadata{} }
func (m *IndexOperationMetadata) String() string { return proto.CompactTextString(m) }
func (*IndexOperationMetadata) ProtoMessage()    {}
func (*IndexOperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{0}
}
func (m *IndexOperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexOperationMetadata.Unmarshal(m, b)
}
func (m *IndexOperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexOperationMetadata.Marshal(b, m, deterministic)
}
func (dst *IndexOperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexOperationMetadata.Merge(dst, src)
}
func (m *IndexOperationMetadata) XXX_Size() int {
	return xxx_messageInfo_IndexOperationMetadata.Size(m)
}
func (m *IndexOperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexOperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_IndexOperationMetadata proto.InternalMessageInfo

func (m *IndexOperationMetadata) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *IndexOperationMetadata) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *IndexOperationMetadata) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *IndexOperationMetadata) GetOperationType() IndexOperationMetadata_OperationType {
	if m != nil {
		return m.OperationType
	}
	return IndexOperationMetadata_OPERATION_TYPE_UNSPECIFIED
}

func (m *IndexOperationMetadata) GetCancelled() bool {
	if m != nil {
		return m.Cancelled
	}
	return false
}

func (m *IndexOperationMetadata) GetDocumentProgress() *Progress {
	if m != nil {
		return m.DocumentProgress
	}
	return nil
}

// Measures the progress of a particular metric.
type Progress struct {
	// An estimate of how much work has been completed. Note that this may be
	// greater than `work_estimated`.
	WorkCompleted int64 `protobuf:"varint,1,opt,name=work_completed,json=workCompleted" json:"work_completed,omitempty"`
	// An estimate of how much work needs to be performed. Zero if the
	// work estimate is unavailable. May change as work progresses.
	WorkEstimated        int64    `protobuf:"varint,2,opt,name=work_estimated,json=workEstimated" json:"work_estimated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Progress) Reset()         { *m = Progress{} }
func (m *Progress) String() string { return proto.CompactTextString(m) }
func (*Progress) ProtoMessage()    {}
func (*Progress) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{1}
}
func (m *Progress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Progress.Unmarshal(m, b)
}
func (m *Progress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Progress.Marshal(b, m, deterministic)
}
func (dst *Progress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Progress.Merge(dst, src)
}
func (m *Progress) XXX_Size() int {
	return xxx_messageInfo_Progress.Size(m)
}
func (m *Progress) XXX_DiscardUnknown() {
	xxx_messageInfo_Progress.DiscardUnknown(m)
}

var xxx_messageInfo_Progress proto.InternalMessageInfo

func (m *Progress) GetWorkCompleted() int64 {
	if m != nil {
		return m.WorkCompleted
	}
	return 0
}

func (m *Progress) GetWorkEstimated() int64 {
	if m != nil {
		return m.WorkEstimated
	}
	return 0
}

// The request for [FirestoreAdmin.CreateIndex][google.firestore.admin.v1beta1.FirestoreAdmin.CreateIndex].
type CreateIndexRequest struct {
	// The name of the database this index will apply to. For example:
	// `projects/{project_id}/databases/{database_id}`
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	// The index to create. The name and state fields are output only and will be
	// ignored. Certain single field indexes cannot be created or deleted.
	Index                *Index   `protobuf:"bytes,2,opt,name=index" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateIndexRequest) Reset()         { *m = CreateIndexRequest{} }
func (m *CreateIndexRequest) String() string { return proto.CompactTextString(m) }
func (*CreateIndexRequest) ProtoMessage()    {}
func (*CreateIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{2}
}
func (m *CreateIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateIndexRequest.Unmarshal(m, b)
}
func (m *CreateIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateIndexRequest.Marshal(b, m, deterministic)
}
func (dst *CreateIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateIndexRequest.Merge(dst, src)
}
func (m *CreateIndexRequest) XXX_Size() int {
	return xxx_messageInfo_CreateIndexRequest.Size(m)
}
func (m *CreateIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateIndexRequest proto.InternalMessageInfo

func (m *CreateIndexRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateIndexRequest) GetIndex() *Index {
	if m != nil {
		return m.Index
	}
	return nil
}

// The request for [FirestoreAdmin.GetIndex][google.firestore.admin.v1beta1.FirestoreAdmin.GetIndex].
type GetIndexRequest struct {
	// The name of the index. For example:
	// `projects/{project_id}/databases/{database_id}/indexes/{index_id}`
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetIndexRequest) Reset()         { *m = GetIndexRequest{} }
func (m *GetIndexRequest) String() string { return proto.CompactTextString(m) }
func (*GetIndexRequest) ProtoMessage()    {}
func (*GetIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{3}
}
func (m *GetIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetIndexRequest.Unmarshal(m, b)
}
func (m *GetIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetIndexRequest.Marshal(b, m, deterministic)
}
func (dst *GetIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetIndexRequest.Merge(dst, src)
}
func (m *GetIndexRequest) XXX_Size() int {
	return xxx_messageInfo_GetIndexRequest.Size(m)
}
func (m *GetIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetIndexRequest proto.InternalMessageInfo

func (m *GetIndexRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The request for [FirestoreAdmin.ListIndexes][google.firestore.admin.v1beta1.FirestoreAdmin.ListIndexes].
type ListIndexesRequest struct {
	// The database name. For example:
	// `projects/{project_id}/databases/{database_id}`
	Parent string `protobuf:"bytes,1,opt,name=parent" json:"parent,omitempty"`
	Filter string `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// The standard List page size.
	PageSize int32 `protobuf:"varint,3,opt,name=page_size,json=pageSize" json:"page_size,omitempty"`
	// The standard List page token.
	PageToken            string   `protobuf:"bytes,4,opt,name=page_token,json=pageToken" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIndexesRequest) Reset()         { *m = ListIndexesRequest{} }
func (m *ListIndexesRequest) String() string { return proto.CompactTextString(m) }
func (*ListIndexesRequest) ProtoMessage()    {}
func (*ListIndexesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{4}
}
func (m *ListIndexesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIndexesRequest.Unmarshal(m, b)
}
func (m *ListIndexesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIndexesRequest.Marshal(b, m, deterministic)
}
func (dst *ListIndexesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIndexesRequest.Merge(dst, src)
}
func (m *ListIndexesRequest) XXX_Size() int {
	return xxx_messageInfo_ListIndexesRequest.Size(m)
}
func (m *ListIndexesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIndexesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListIndexesRequest proto.InternalMessageInfo

func (m *ListIndexesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListIndexesRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

func (m *ListIndexesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListIndexesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// The request for [FirestoreAdmin.DeleteIndex][google.firestore.admin.v1beta1.FirestoreAdmin.DeleteIndex].
type DeleteIndexRequest struct {
	// The index name. For example:
	// `projects/{project_id}/databases/{database_id}/indexes/{index_id}`
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteIndexRequest) Reset()         { *m = DeleteIndexRequest{} }
func (m *DeleteIndexRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteIndexRequest) ProtoMessage()    {}
func (*DeleteIndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{5}
}
func (m *DeleteIndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteIndexRequest.Unmarshal(m, b)
}
func (m *DeleteIndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteIndexRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteIndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteIndexRequest.Merge(dst, src)
}
func (m *DeleteIndexRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteIndexRequest.Size(m)
}
func (m *DeleteIndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteIndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteIndexRequest proto.InternalMessageInfo

func (m *DeleteIndexRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response for [FirestoreAdmin.ListIndexes][google.firestore.admin.v1beta1.FirestoreAdmin.ListIndexes].
type ListIndexesResponse struct {
	// The indexes.
	Indexes []*Index `protobuf:"bytes,1,rep,name=indexes" json:"indexes,omitempty"`
	// The standard List next-page token.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListIndexesResponse) Reset()         { *m = ListIndexesResponse{} }
func (m *ListIndexesResponse) String() string { return proto.CompactTextString(m) }
func (*ListIndexesResponse) ProtoMessage()    {}
func (*ListIndexesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_firestore_admin_bf8d6f204477dd80, []int{6}
}
func (m *ListIndexesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListIndexesResponse.Unmarshal(m, b)
}
func (m *ListIndexesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListIndexesResponse.Marshal(b, m, deterministic)
}
func (dst *ListIndexesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListIndexesResponse.Merge(dst, src)
}
func (m *ListIndexesResponse) XXX_Size() int {
	return xxx_messageInfo_ListIndexesResponse.Size(m)
}
func (m *ListIndexesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListIndexesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListIndexesResponse proto.InternalMessageInfo

func (m *ListIndexesResponse) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *ListIndexesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*IndexOperationMetadata)(nil), "google.firestore.admin.v1beta1.IndexOperationMetadata")
	proto.RegisterType((*Progress)(nil), "google.firestore.admin.v1beta1.Progress")
	proto.RegisterType((*CreateIndexRequest)(nil), "google.firestore.admin.v1beta1.CreateIndexRequest")
	proto.RegisterType((*GetIndexRequest)(nil), "google.firestore.admin.v1beta1.GetIndexRequest")
	proto.RegisterType((*ListIndexesRequest)(nil), "google.firestore.admin.v1beta1.ListIndexesRequest")
	proto.RegisterType((*DeleteIndexRequest)(nil), "google.firestore.admin.v1beta1.DeleteIndexRequest")
	proto.RegisterType((*ListIndexesResponse)(nil), "google.firestore.admin.v1beta1.ListIndexesResponse")
	proto.RegisterEnum("google.firestore.admin.v1beta1.IndexOperationMetadata_OperationType", IndexOperationMetadata_OperationType_name, IndexOperationMetadata_OperationType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FirestoreAdminClient is the client API for FirestoreAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FirestoreAdminClient interface {
	// Creates the specified index.
	// A newly created index's initial state is `CREATING`. On completion of the
	// returned [google.longrunning.Operation][google.longrunning.Operation], the state will be `READY`.
	// If the index already exists, the call will return an `ALREADY_EXISTS`
	// status.
	//
	// During creation, the process could result in an error, in which case the
	// index will move to the `ERROR` state. The process can be recovered by
	// fixing the data that caused the error, removing the index with
	// [delete][google.firestore.admin.v1beta1.FirestoreAdmin.DeleteIndex], then re-creating the index with
	// [create][google.firestore.admin.v1beta1.FirestoreAdmin.CreateIndex].
	//
	// Indexes with a single field cannot be created.
	CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*longrunning.Operation, error)
	// Lists the indexes that match the specified filters.
	ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error)
	// Gets an index.
	GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*Index, error)
	// Deletes an index.
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type firestoreAdminClient struct {
	cc *grpc.ClientConn
}

func NewFirestoreAdminClient(cc *grpc.ClientConn) FirestoreAdminClient {
	return &firestoreAdminClient{cc}
}

func (c *firestoreAdminClient) CreateIndex(ctx context.Context, in *CreateIndexRequest, opts ...grpc.CallOption) (*longrunning.Operation, error) {
	out := new(longrunning.Operation)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta1.FirestoreAdmin/CreateIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firestoreAdminClient) ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error) {
	out := new(ListIndexesResponse)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta1.FirestoreAdmin/ListIndexes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firestoreAdminClient) GetIndex(ctx context.Context, in *GetIndexRequest, opts ...grpc.CallOption) (*Index, error) {
	out := new(Index)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta1.FirestoreAdmin/GetIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *firestoreAdminClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.firestore.admin.v1beta1.FirestoreAdmin/DeleteIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FirestoreAdmin service

type FirestoreAdminServer interface {
	// Creates the specified index.
	// A newly created index's initial state is `CREATING`. On completion of the
	// returned [google.longrunning.Operation][google.longrunning.Operation], the state will be `READY`.
	// If the index already exists, the call will return an `ALREADY_EXISTS`
	// status.
	//
	// During creation, the process could result in an error, in which case the
	// index will move to the `ERROR` state. The process can be recovered by
	// fixing the data that caused the error, removing the index with
	// [delete][google.firestore.admin.v1beta1.FirestoreAdmin.DeleteIndex], then re-creating the index with
	// [create][google.firestore.admin.v1beta1.FirestoreAdmin.CreateIndex].
	//
	// Indexes with a single field cannot be created.
	CreateIndex(context.Context, *CreateIndexRequest) (*longrunning.Operation, error)
	// Lists the indexes that match the specified filters.
	ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error)
	// Gets an index.
	GetIndex(context.Context, *GetIndexRequest) (*Index, error)
	// Deletes an index.
	DeleteIndex(context.Context, *DeleteIndexRequest) (*empty.Empty, error)
}

func RegisterFirestoreAdminServer(s *grpc.Server, srv FirestoreAdminServer) {
	s.RegisterService(&_FirestoreAdmin_serviceDesc, srv)
}

func _FirestoreAdmin_CreateIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).CreateIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta1.FirestoreAdmin/CreateIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).CreateIndex(ctx, req.(*CreateIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirestoreAdmin_ListIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).ListIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta1.FirestoreAdmin/ListIndexes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).ListIndexes(ctx, req.(*ListIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirestoreAdmin_GetIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).GetIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta1.FirestoreAdmin/GetIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).GetIndex(ctx, req.(*GetIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FirestoreAdmin_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FirestoreAdminServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.firestore.admin.v1beta1.FirestoreAdmin/DeleteIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FirestoreAdminServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FirestoreAdmin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.firestore.admin.v1beta1.FirestoreAdmin",
	HandlerType: (*FirestoreAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIndex",
			Handler:    _FirestoreAdmin_CreateIndex_Handler,
		},
		{
			MethodName: "ListIndexes",
			Handler:    _FirestoreAdmin_ListIndexes_Handler,
		},
		{
			MethodName: "GetIndex",
			Handler:    _FirestoreAdmin_GetIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _FirestoreAdmin_DeleteIndex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/firestore/admin/v1beta1/firestore_admin.proto",
}

func init() {
	proto.RegisterFile("google/firestore/admin/v1beta1/firestore_admin.proto", fileDescriptor_firestore_admin_bf8d6f204477dd80)
}

var fileDescriptor_firestore_admin_bf8d6f204477dd80 = []byte{
	// 841 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0xe9, 0xcf, 0x26, 0xa7, 0x6a, 0xb6, 0xcc, 0xa2, 0x2a, 0xf2, 0xfe, 0x10, 0x19, 0x8a,
	0xa2, 0x5c, 0xd8, 0x34, 0x0b, 0x12, 0xcb, 0x0a, 0xad, 0x5a, 0xd7, 0xad, 0x22, 0x41, 0x1b, 0xb9,
	0x59, 0xb4, 0x70, 0x63, 0x4d, 0xe3, 0x53, 0xcb, 0xd4, 0x9e, 0x31, 0x9e, 0x09, 0x6c, 0x17, 0x2d,
	0x42, 0xbc, 0xc2, 0xde, 0xee, 0x0d, 0x5c, 0x72, 0x81, 0x78, 0x0b, 0x1e, 0x80, 0x57, 0xe0, 0x41,
	0x90, 0xc7, 0x9e, 0x34, 0xd9, 0x6e, 0x71, 0x7b, 0x97, 0xf3, 0xcd, 0xf9, 0xbe, 0xf3, 0x9d, 0x39,
	0xc7, 0x13, 0xf8, 0x24, 0xe2, 0x3c, 0x4a, 0xd0, 0x39, 0x8d, 0x73, 0x14, 0x92, 0xe7, 0xe8, 0xd0,
	0x30, 0x8d, 0x99, 0xf3, 0xc3, 0xf6, 0x09, 0x4a, 0xba, 0x7d, 0x81, 0x07, 0x0a, 0xb7, 0xb3, 0x9c,
	0x4b, 0x4e, 0x1e, 0x94, 0x2c, 0x7b, 0x76, 0x6a, 0x97, 0xa7, 0x15, 0xcb, 0xbc, 0x57, 0xa9, 0xd2,
	0x2c, 0x76, 0x28, 0x63, 0x5c, 0x52, 0x19, 0x73, 0x26, 0x4a, 0xb6, 0xd9, 0xaf, 0xa9, 0x19, 0xb3,
	0x10, 0x9f, 0x57, 0xb9, 0x1f, 0x54, 0xb9, 0x09, 0x67, 0x51, 0x3e, 0x65, 0x2c, 0x66, 0x91, 0xc3,
	0x33, 0xcc, 0x17, 0x04, 0xef, 0x56, 0x49, 0x2a, 0x3a, 0x99, 0x9e, 0x3a, 0x98, 0x66, 0xf2, 0xbc,
	0x3a, 0x7c, 0xff, 0xcd, 0x43, 0x19, 0xa7, 0x28, 0x24, 0x4d, 0xb3, 0x32, 0xc1, 0xfa, 0x7b, 0x09,
	0x36, 0x87, 0x45, 0xc9, 0x23, 0xad, 0xfb, 0x15, 0x4a, 0x1a, 0x52, 0x49, 0xc9, 0x23, 0x00, 0x21,
	0x69, 0x2e, 0x83, 0x82, 0xd3, 0x31, 0xba, 0x46, 0x6f, 0x6d, 0x60, 0xda, 0x55, 0xf3, 0x5a, 0xd0,
	0x1e, 0x6b, 0x41, 0xbf, 0xa5, 0xb2, 0x8b, 0x98, 0x7c, 0x0a, 0x4d, 0x64, 0x61, 0x49, 0x6c, 0xd4,
	0x12, 0x6f, 0x21, 0x0b, 0x15, 0xed, 0x3d, 0x58, 0x51, 0xed, 0x77, 0x96, 0xba, 0x46, 0xaf, 0xe5,
	0x97, 0x01, 0x39, 0x83, 0xf6, 0xac, 0xe9, 0x40, 0x9e, 0x67, 0xd8, 0x59, 0xee, 0x1a, 0xbd, 0xf6,
	0x60, 0xcf, 0xfe, 0xff, 0x41, 0xd8, 0x6f, 0xef, 0xcb, 0x9e, 0x21, 0xe3, 0xf3, 0x0c, 0xfd, 0x75,
	0x3e, 0x1f, 0x92, 0x7b, 0xd0, 0x9a, 0x50, 0x36, 0xc1, 0x24, 0xc1, 0xb0, 0xb3, 0xd2, 0x35, 0x7a,
	0x4d, 0xff, 0x02, 0x20, 0x4f, 0xe1, 0xdd, 0x90, 0x4f, 0xa6, 0x29, 0x32, 0x19, 0x64, 0x39, 0x8f,
	0x72, 0x14, 0xa2, 0xb3, 0xaa, 0x1a, 0xec, 0xd5, 0xb9, 0x19, 0x55, 0xf9, 0xfe, 0x86, 0x96, 0xd0,
	0x88, 0xe5, 0xc2, 0xfa, 0x82, 0x29, 0xf2, 0x00, 0xcc, 0xa3, 0x91, 0xe7, 0xef, 0x8c, 0x87, 0x47,
	0x87, 0xc1, 0xf8, 0x9b, 0x91, 0x17, 0x3c, 0x3d, 0x3c, 0x1e, 0x79, 0xee, 0x70, 0x7f, 0xe8, 0xed,
	0x6d, 0xbc, 0x43, 0x08, 0xb4, 0x5d, 0xdf, 0xdb, 0x19, 0x0f, 0x0f, 0x0f, 0x82, 0xe1, 0xe1, 0x9e,
	0xf7, 0x6c, 0xc3, 0xb0, 0x9e, 0x41, 0x53, 0x0b, 0x92, 0x2d, 0x68, 0xff, 0xc8, 0xf3, 0xb3, 0x60,
	0xc2, 0xd3, 0x2c, 0x41, 0x89, 0xa1, 0x1a, 0xdf, 0x92, 0xbf, 0x5e, 0xa0, 0xae, 0x06, 0x67, 0x69,
	0x28, 0x64, 0x9c, 0xd2, 0x22, 0xad, 0x71, 0x91, 0xe6, 0x69, 0xd0, 0x8a, 0x81, 0xb8, 0x39, 0x52,
	0x89, 0xea, 0x42, 0x7d, 0xfc, 0x7e, 0x8a, 0x42, 0x92, 0x4d, 0x58, 0xcd, 0x68, 0x8e, 0x4c, 0x2a,
	0xed, 0x96, 0x5f, 0x45, 0xe4, 0xb1, 0x1e, 0x62, 0x39, 0xf8, 0xad, 0x6b, 0x4d, 0xa9, 0x9a, 0xb5,
	0xb5, 0x05, 0xb7, 0x0f, 0x50, 0x2e, 0xd4, 0x21, 0xb0, 0xcc, 0x68, 0xb5, 0x80, 0x2d, 0x5f, 0xfd,
	0xb6, 0x7e, 0x31, 0x80, 0x7c, 0x19, 0x8b, 0x32, 0x11, 0x45, 0x9d, 0xa5, 0x4d, 0x58, 0x3d, 0x8d,
	0x13, 0x89, 0xb9, 0xf2, 0xd4, 0xf2, 0xab, 0x88, 0xdc, 0x85, 0x56, 0x46, 0x23, 0x0c, 0x44, 0xfc,
	0x02, 0xd5, 0xce, 0xad, 0xf8, 0xcd, 0x02, 0x38, 0x8e, 0x5f, 0x20, 0xb9, 0x0f, 0xa0, 0x0e, 0x25,
	0x3f, 0x43, 0xa6, 0x56, 0xae, 0xe5, 0xab, 0xf4, 0x71, 0x01, 0x58, 0x3d, 0x20, 0x7b, 0x58, 0x5c,
	0x63, 0xad, 0xd9, 0x9f, 0xe1, 0xce, 0x82, 0x57, 0x91, 0x71, 0x26, 0x90, 0x3c, 0x81, 0x5b, 0x71,
	0x09, 0x75, 0x8c, 0xee, 0xd2, 0xf5, 0x6f, 0x4a, 0xb3, 0xc8, 0x47, 0x70, 0x9b, 0xe1, 0x73, 0x19,
	0xcc, 0xb9, 0x2c, 0xdb, 0x5b, 0x2f, 0xe0, 0x91, 0x76, 0x3a, 0x78, 0xbd, 0x02, 0xed, 0x7d, 0x2d,
	0xb9, 0x53, 0x28, 0x92, 0xdf, 0x0c, 0x58, 0x9b, 0x1b, 0x29, 0x19, 0xd4, 0x95, 0xbe, 0x3c, 0x7f,
	0xf3, 0xbe, 0xe6, 0xcc, 0xbd, 0x4e, 0x17, 0xdf, 0x96, 0xf5, 0xe4, 0xd7, 0x7f, 0xfe, 0x7d, 0xd5,
	0x78, 0x64, 0x7d, 0x3c, 0x7b, 0xd9, 0x7e, 0x2a, 0xa7, 0xf1, 0x45, 0x96, 0xf3, 0xef, 0x70, 0x22,
	0x85, 0xd3, 0x77, 0x8a, 0xef, 0xf1, 0x84, 0x0a, 0x14, 0x4e, 0xff, 0xa5, 0x53, 0xf5, 0xf5, 0x79,
	0xf5, 0xd9, 0xff, 0x65, 0xc0, 0xda, 0xdc, 0xbd, 0xd5, 0x7b, 0xbc, 0xbc, 0x10, 0xe6, 0xc3, 0x1b,
	0x71, 0xca, 0xc1, 0x58, 0x9f, 0x29, 0xe7, 0x03, 0x72, 0x63, 0xe7, 0xe4, 0xb5, 0x01, 0x4d, 0xbd,
	0xbe, 0xc4, 0xa9, 0xab, 0xfd, 0xc6, 0xa2, 0x9b, 0xd7, 0x9b, 0xff, 0xdb, 0xec, 0x15, 0x6b, 0x76,
	0x85, 0x39, 0xed, 0xcd, 0xe9, 0xbf, 0x24, 0xaf, 0x0c, 0x58, 0x9b, 0xdb, 0xd9, 0xfa, 0x1b, 0xbd,
	0xbc, 0xe0, 0xe6, 0xe6, 0xa5, 0x77, 0xdc, 0x2b, 0xfe, 0x6e, 0xb4, 0xab, 0xfe, 0x8d, 0x5d, 0xed,
	0xfe, 0x69, 0x80, 0x35, 0xe1, 0x69, 0x8d, 0x97, 0xdd, 0x3b, 0x8b, 0x2b, 0x3c, 0x2a, 0xca, 0x8f,
	0x8c, 0x6f, 0xdd, 0x8a, 0x16, 0xf1, 0x84, 0xb2, 0xc8, 0xe6, 0x79, 0xe4, 0x44, 0xc8, 0x94, 0x39,
	0xa7, 0x3c, 0xa2, 0x59, 0x2c, 0xae, 0xfa, 0xb7, 0x7d, 0xac, 0xa2, 0xdf, 0x1b, 0xcb, 0x07, 0xee,
	0xfe, 0xf1, 0x1f, 0x8d, 0x0f, 0x0f, 0x4a, 0x31, 0x37, 0xe1, 0xd3, 0xd0, 0x9e, 0x15, 0xb4, 0x55,
	0x45, 0xfb, 0xeb, 0xed, 0xdd, 0x82, 0x73, 0xb2, 0xaa, 0xd4, 0x1f, 0xfe, 0x17, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0x78, 0x44, 0x07, 0x3e, 0x08, 0x00, 0x00,
}
