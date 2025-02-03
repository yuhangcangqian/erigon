// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: execution/execution.proto

package executionproto

import (
	context "context"
	typesproto "github.com/erigontech/erigon-lib/gointerfaces/typesproto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Execution_InsertBlocks_FullMethodName        = "/execution.Execution/InsertBlocks"
	Execution_ValidateChain_FullMethodName       = "/execution.Execution/ValidateChain"
	Execution_UpdateForkChoice_FullMethodName    = "/execution.Execution/UpdateForkChoice"
	Execution_AssembleBlock_FullMethodName       = "/execution.Execution/AssembleBlock"
	Execution_GetAssembledBlock_FullMethodName   = "/execution.Execution/GetAssembledBlock"
	Execution_CurrentHeader_FullMethodName       = "/execution.Execution/CurrentHeader"
	Execution_GetTD_FullMethodName               = "/execution.Execution/GetTD"
	Execution_GetHeader_FullMethodName           = "/execution.Execution/GetHeader"
	Execution_GetBody_FullMethodName             = "/execution.Execution/GetBody"
	Execution_HasBlock_FullMethodName            = "/execution.Execution/HasBlock"
	Execution_GetBodiesByRange_FullMethodName    = "/execution.Execution/GetBodiesByRange"
	Execution_GetBodiesByHashes_FullMethodName   = "/execution.Execution/GetBodiesByHashes"
	Execution_IsCanonicalHash_FullMethodName     = "/execution.Execution/IsCanonicalHash"
	Execution_GetHeaderHashNumber_FullMethodName = "/execution.Execution/GetHeaderHashNumber"
	Execution_GetForkChoice_FullMethodName       = "/execution.Execution/GetForkChoice"
	Execution_Ready_FullMethodName               = "/execution.Execution/Ready"
	Execution_FrozenBlocks_FullMethodName        = "/execution.Execution/FrozenBlocks"
)

// ExecutionClient is the client API for Execution service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutionClient interface {
	// Chain Putters.
	InsertBlocks(ctx context.Context, in *InsertBlocksRequest, opts ...grpc.CallOption) (*InsertionResult, error)
	// Chain Validation and ForkChoice.
	ValidateChain(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationReceipt, error)
	UpdateForkChoice(ctx context.Context, in *ForkChoice, opts ...grpc.CallOption) (*ForkChoiceReceipt, error)
	// Block Assembly
	// EAGAIN design here, AssembleBlock initiates the asynchronous request, and GetAssembleBlock just return it if ready.
	AssembleBlock(ctx context.Context, in *AssembleBlockRequest, opts ...grpc.CallOption) (*AssembleBlockResponse, error)
	GetAssembledBlock(ctx context.Context, in *GetAssembledBlockRequest, opts ...grpc.CallOption) (*GetAssembledBlockResponse, error)
	// Chain Getters.
	CurrentHeader(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHeaderResponse, error)
	GetTD(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetTDResponse, error)
	GetHeader(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetHeaderResponse, error)
	GetBody(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetBodyResponse, error)
	HasBlock(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*HasBlockResponse, error)
	// Ranges
	GetBodiesByRange(ctx context.Context, in *GetBodiesByRangeRequest, opts ...grpc.CallOption) (*GetBodiesBatchResponse, error)
	GetBodiesByHashes(ctx context.Context, in *GetBodiesByHashesRequest, opts ...grpc.CallOption) (*GetBodiesBatchResponse, error)
	// Chain checkers
	IsCanonicalHash(ctx context.Context, in *typesproto.H256, opts ...grpc.CallOption) (*IsCanonicalResponse, error)
	GetHeaderHashNumber(ctx context.Context, in *typesproto.H256, opts ...grpc.CallOption) (*GetHeaderHashNumberResponse, error)
	GetForkChoice(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ForkChoice, error)
	// Misc
	// We want to figure out whether we processed snapshots and cleanup sync cycles.
	Ready(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReadyResponse, error)
	// Frozen blocks are how many blocks are in snapshots .seg files.
	FrozenBlocks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FrozenBlocksResponse, error)
}

type executionClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutionClient(cc grpc.ClientConnInterface) ExecutionClient {
	return &executionClient{cc}
}

func (c *executionClient) InsertBlocks(ctx context.Context, in *InsertBlocksRequest, opts ...grpc.CallOption) (*InsertionResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertionResult)
	err := c.cc.Invoke(ctx, Execution_InsertBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) ValidateChain(ctx context.Context, in *ValidationRequest, opts ...grpc.CallOption) (*ValidationReceipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidationReceipt)
	err := c.cc.Invoke(ctx, Execution_ValidateChain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) UpdateForkChoice(ctx context.Context, in *ForkChoice, opts ...grpc.CallOption) (*ForkChoiceReceipt, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForkChoiceReceipt)
	err := c.cc.Invoke(ctx, Execution_UpdateForkChoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) AssembleBlock(ctx context.Context, in *AssembleBlockRequest, opts ...grpc.CallOption) (*AssembleBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssembleBlockResponse)
	err := c.cc.Invoke(ctx, Execution_AssembleBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetAssembledBlock(ctx context.Context, in *GetAssembledBlockRequest, opts ...grpc.CallOption) (*GetAssembledBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAssembledBlockResponse)
	err := c.cc.Invoke(ctx, Execution_GetAssembledBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) CurrentHeader(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetHeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHeaderResponse)
	err := c.cc.Invoke(ctx, Execution_CurrentHeader_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetTD(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetTDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTDResponse)
	err := c.cc.Invoke(ctx, Execution_GetTD_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetHeader(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetHeaderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHeaderResponse)
	err := c.cc.Invoke(ctx, Execution_GetHeader_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetBody(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*GetBodyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBodyResponse)
	err := c.cc.Invoke(ctx, Execution_GetBody_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) HasBlock(ctx context.Context, in *GetSegmentRequest, opts ...grpc.CallOption) (*HasBlockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HasBlockResponse)
	err := c.cc.Invoke(ctx, Execution_HasBlock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetBodiesByRange(ctx context.Context, in *GetBodiesByRangeRequest, opts ...grpc.CallOption) (*GetBodiesBatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBodiesBatchResponse)
	err := c.cc.Invoke(ctx, Execution_GetBodiesByRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetBodiesByHashes(ctx context.Context, in *GetBodiesByHashesRequest, opts ...grpc.CallOption) (*GetBodiesBatchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBodiesBatchResponse)
	err := c.cc.Invoke(ctx, Execution_GetBodiesByHashes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) IsCanonicalHash(ctx context.Context, in *typesproto.H256, opts ...grpc.CallOption) (*IsCanonicalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IsCanonicalResponse)
	err := c.cc.Invoke(ctx, Execution_IsCanonicalHash_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetHeaderHashNumber(ctx context.Context, in *typesproto.H256, opts ...grpc.CallOption) (*GetHeaderHashNumberResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHeaderHashNumberResponse)
	err := c.cc.Invoke(ctx, Execution_GetHeaderHashNumber_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) GetForkChoice(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ForkChoice, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ForkChoice)
	err := c.cc.Invoke(ctx, Execution_GetForkChoice_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) Ready(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ReadyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadyResponse)
	err := c.cc.Invoke(ctx, Execution_Ready_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executionClient) FrozenBlocks(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*FrozenBlocksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FrozenBlocksResponse)
	err := c.cc.Invoke(ctx, Execution_FrozenBlocks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutionServer is the server API for Execution service.
// All implementations must embed UnimplementedExecutionServer
// for forward compatibility.
type ExecutionServer interface {
	// Chain Putters.
	InsertBlocks(context.Context, *InsertBlocksRequest) (*InsertionResult, error)
	// Chain Validation and ForkChoice.
	ValidateChain(context.Context, *ValidationRequest) (*ValidationReceipt, error)
	UpdateForkChoice(context.Context, *ForkChoice) (*ForkChoiceReceipt, error)
	// Block Assembly
	// EAGAIN design here, AssembleBlock initiates the asynchronous request, and GetAssembleBlock just return it if ready.
	AssembleBlock(context.Context, *AssembleBlockRequest) (*AssembleBlockResponse, error)
	GetAssembledBlock(context.Context, *GetAssembledBlockRequest) (*GetAssembledBlockResponse, error)
	// Chain Getters.
	CurrentHeader(context.Context, *emptypb.Empty) (*GetHeaderResponse, error)
	GetTD(context.Context, *GetSegmentRequest) (*GetTDResponse, error)
	GetHeader(context.Context, *GetSegmentRequest) (*GetHeaderResponse, error)
	GetBody(context.Context, *GetSegmentRequest) (*GetBodyResponse, error)
	HasBlock(context.Context, *GetSegmentRequest) (*HasBlockResponse, error)
	// Ranges
	GetBodiesByRange(context.Context, *GetBodiesByRangeRequest) (*GetBodiesBatchResponse, error)
	GetBodiesByHashes(context.Context, *GetBodiesByHashesRequest) (*GetBodiesBatchResponse, error)
	// Chain checkers
	IsCanonicalHash(context.Context, *typesproto.H256) (*IsCanonicalResponse, error)
	GetHeaderHashNumber(context.Context, *typesproto.H256) (*GetHeaderHashNumberResponse, error)
	GetForkChoice(context.Context, *emptypb.Empty) (*ForkChoice, error)
	// Misc
	// We want to figure out whether we processed snapshots and cleanup sync cycles.
	Ready(context.Context, *emptypb.Empty) (*ReadyResponse, error)
	// Frozen blocks are how many blocks are in snapshots .seg files.
	FrozenBlocks(context.Context, *emptypb.Empty) (*FrozenBlocksResponse, error)
	mustEmbedUnimplementedExecutionServer()
}

// UnimplementedExecutionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExecutionServer struct{}

func (UnimplementedExecutionServer) InsertBlocks(context.Context, *InsertBlocksRequest) (*InsertionResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertBlocks not implemented")
}
func (UnimplementedExecutionServer) ValidateChain(context.Context, *ValidationRequest) (*ValidationReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateChain not implemented")
}
func (UnimplementedExecutionServer) UpdateForkChoice(context.Context, *ForkChoice) (*ForkChoiceReceipt, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateForkChoice not implemented")
}
func (UnimplementedExecutionServer) AssembleBlock(context.Context, *AssembleBlockRequest) (*AssembleBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssembleBlock not implemented")
}
func (UnimplementedExecutionServer) GetAssembledBlock(context.Context, *GetAssembledBlockRequest) (*GetAssembledBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssembledBlock not implemented")
}
func (UnimplementedExecutionServer) CurrentHeader(context.Context, *emptypb.Empty) (*GetHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentHeader not implemented")
}
func (UnimplementedExecutionServer) GetTD(context.Context, *GetSegmentRequest) (*GetTDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTD not implemented")
}
func (UnimplementedExecutionServer) GetHeader(context.Context, *GetSegmentRequest) (*GetHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeader not implemented")
}
func (UnimplementedExecutionServer) GetBody(context.Context, *GetSegmentRequest) (*GetBodyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBody not implemented")
}
func (UnimplementedExecutionServer) HasBlock(context.Context, *GetSegmentRequest) (*HasBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasBlock not implemented")
}
func (UnimplementedExecutionServer) GetBodiesByRange(context.Context, *GetBodiesByRangeRequest) (*GetBodiesBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBodiesByRange not implemented")
}
func (UnimplementedExecutionServer) GetBodiesByHashes(context.Context, *GetBodiesByHashesRequest) (*GetBodiesBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBodiesByHashes not implemented")
}
func (UnimplementedExecutionServer) IsCanonicalHash(context.Context, *typesproto.H256) (*IsCanonicalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCanonicalHash not implemented")
}
func (UnimplementedExecutionServer) GetHeaderHashNumber(context.Context, *typesproto.H256) (*GetHeaderHashNumberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeaderHashNumber not implemented")
}
func (UnimplementedExecutionServer) GetForkChoice(context.Context, *emptypb.Empty) (*ForkChoice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForkChoice not implemented")
}
func (UnimplementedExecutionServer) Ready(context.Context, *emptypb.Empty) (*ReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ready not implemented")
}
func (UnimplementedExecutionServer) FrozenBlocks(context.Context, *emptypb.Empty) (*FrozenBlocksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FrozenBlocks not implemented")
}
func (UnimplementedExecutionServer) mustEmbedUnimplementedExecutionServer() {}
func (UnimplementedExecutionServer) testEmbeddedByValue()                   {}

// UnsafeExecutionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutionServer will
// result in compilation errors.
type UnsafeExecutionServer interface {
	mustEmbedUnimplementedExecutionServer()
}

func RegisterExecutionServer(s grpc.ServiceRegistrar, srv ExecutionServer) {
	// If the following call pancis, it indicates UnimplementedExecutionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Execution_ServiceDesc, srv)
}

func _Execution_InsertBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertBlocksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).InsertBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_InsertBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).InsertBlocks(ctx, req.(*InsertBlocksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_ValidateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).ValidateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_ValidateChain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).ValidateChain(ctx, req.(*ValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_UpdateForkChoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForkChoice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).UpdateForkChoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_UpdateForkChoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).UpdateForkChoice(ctx, req.(*ForkChoice))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_AssembleBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssembleBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).AssembleBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_AssembleBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).AssembleBlock(ctx, req.(*AssembleBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetAssembledBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAssembledBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetAssembledBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetAssembledBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetAssembledBlock(ctx, req.(*GetAssembledBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_CurrentHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).CurrentHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_CurrentHeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).CurrentHeader(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetTD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetTD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetTD_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetTD(ctx, req.(*GetSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetHeader_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetHeader(ctx, req.(*GetSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetBody(ctx, req.(*GetSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_HasBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSegmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).HasBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_HasBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).HasBlock(ctx, req.(*GetSegmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetBodiesByRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBodiesByRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetBodiesByRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetBodiesByRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetBodiesByRange(ctx, req.(*GetBodiesByRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetBodiesByHashes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBodiesByHashesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetBodiesByHashes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetBodiesByHashes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetBodiesByHashes(ctx, req.(*GetBodiesByHashesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_IsCanonicalHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(typesproto.H256)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).IsCanonicalHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_IsCanonicalHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).IsCanonicalHash(ctx, req.(*typesproto.H256))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetHeaderHashNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(typesproto.H256)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetHeaderHashNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetHeaderHashNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetHeaderHashNumber(ctx, req.(*typesproto.H256))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_GetForkChoice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).GetForkChoice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_GetForkChoice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).GetForkChoice(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_Ready_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).Ready(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_Ready_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).Ready(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Execution_FrozenBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutionServer).FrozenBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Execution_FrozenBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutionServer).FrozenBlocks(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Execution_ServiceDesc is the grpc.ServiceDesc for Execution service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Execution_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "execution.Execution",
	HandlerType: (*ExecutionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertBlocks",
			Handler:    _Execution_InsertBlocks_Handler,
		},
		{
			MethodName: "ValidateChain",
			Handler:    _Execution_ValidateChain_Handler,
		},
		{
			MethodName: "UpdateForkChoice",
			Handler:    _Execution_UpdateForkChoice_Handler,
		},
		{
			MethodName: "AssembleBlock",
			Handler:    _Execution_AssembleBlock_Handler,
		},
		{
			MethodName: "GetAssembledBlock",
			Handler:    _Execution_GetAssembledBlock_Handler,
		},
		{
			MethodName: "CurrentHeader",
			Handler:    _Execution_CurrentHeader_Handler,
		},
		{
			MethodName: "GetTD",
			Handler:    _Execution_GetTD_Handler,
		},
		{
			MethodName: "GetHeader",
			Handler:    _Execution_GetHeader_Handler,
		},
		{
			MethodName: "GetBody",
			Handler:    _Execution_GetBody_Handler,
		},
		{
			MethodName: "HasBlock",
			Handler:    _Execution_HasBlock_Handler,
		},
		{
			MethodName: "GetBodiesByRange",
			Handler:    _Execution_GetBodiesByRange_Handler,
		},
		{
			MethodName: "GetBodiesByHashes",
			Handler:    _Execution_GetBodiesByHashes_Handler,
		},
		{
			MethodName: "IsCanonicalHash",
			Handler:    _Execution_IsCanonicalHash_Handler,
		},
		{
			MethodName: "GetHeaderHashNumber",
			Handler:    _Execution_GetHeaderHashNumber_Handler,
		},
		{
			MethodName: "GetForkChoice",
			Handler:    _Execution_GetForkChoice_Handler,
		},
		{
			MethodName: "Ready",
			Handler:    _Execution_Ready_Handler,
		},
		{
			MethodName: "FrozenBlocks",
			Handler:    _Execution_FrozenBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "execution/execution.proto",
}
