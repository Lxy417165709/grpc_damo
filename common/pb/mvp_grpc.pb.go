// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MvpClient is the client API for Mvp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MvpClient interface {
	GetAllGoodClasses(ctx context.Context, in *GetAllGoodClassesReq, opts ...grpc.CallOption) (*GetAllGoodClassesRes, error)
	GetAllDesks(ctx context.Context, in *GetAllDesksReq, opts ...grpc.CallOption) (*GetAllDesksRes, error)
	GetAllGoods(ctx context.Context, in *GetAllGoodsReq, opts ...grpc.CallOption) (*GetAllGoodsRes, error)
	GetAllGoodOptions(ctx context.Context, in *GetAllGoodOptionsReq, opts ...grpc.CallOption) (*GetAllGoodOptionsRes, error)
	AddGoodClass(ctx context.Context, in *AddGoodClassReq, opts ...grpc.CallOption) (*AddGoodClassRes, error)
	AddDeskClass(ctx context.Context, in *AddDeskClassReq, opts ...grpc.CallOption) (*AddDeskClassRes, error)
	AddGood(ctx context.Context, in *AddGoodReq, opts ...grpc.CallOption) (*AddGoodRes, error)
	CancelGood(ctx context.Context, in *CancelGoodReq, opts ...grpc.CallOption) (*CancelGoodRes, error)
	AddElement(ctx context.Context, in *AddElementReq, opts ...grpc.CallOption) (*AddElementRes, error)
	AddSpace(ctx context.Context, in *AddSpaceReq, opts ...grpc.CallOption) (*AddSpaceRes, error)
	AddDesk(ctx context.Context, in *AddDeskReq, opts ...grpc.CallOption) (*AddDeskRes, error)
	OrderGood(ctx context.Context, in *OrderGoodReq, opts ...grpc.CallOption) (*OrderGoodRes, error)
	OrderDesk(ctx context.Context, in *OrderDeskReq, opts ...grpc.CallOption) (*OrderDeskRes, error)
	ChangeDesk(ctx context.Context, in *ChangeDeskReq, opts ...grpc.CallOption) (*ChangeDeskRes, error)
	CloseDesk(ctx context.Context, in *CloseDeskReq, opts ...grpc.CallOption) (*CloseDeskRes, error)
	GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderRes, error)
	CheckOut(ctx context.Context, in *CheckOutReq, opts ...grpc.CallOption) (*CheckOutRes, error)
	AddFavorForGood(ctx context.Context, in *AddFavorForGoodReq, opts ...grpc.CallOption) (*AddFavorForGoodRes, error)
	DeleteFavorForGood(ctx context.Context, in *DeleteFavorForGoodReq, opts ...grpc.CallOption) (*DeleteFavorForGoodRes, error)
	GetAllDeskClasses(ctx context.Context, in *GetAllDeskClassesReq, opts ...grpc.CallOption) (*GetAllDeskClassesRes, error)
}

type mvpClient struct {
	cc grpc.ClientConnInterface
}

func NewMvpClient(cc grpc.ClientConnInterface) MvpClient {
	return &mvpClient{cc}
}

func (c *mvpClient) GetAllGoodClasses(ctx context.Context, in *GetAllGoodClassesReq, opts ...grpc.CallOption) (*GetAllGoodClassesRes, error) {
	out := new(GetAllGoodClassesRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetAllGoodClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetAllDesks(ctx context.Context, in *GetAllDesksReq, opts ...grpc.CallOption) (*GetAllDesksRes, error) {
	out := new(GetAllDesksRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetAllDesks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetAllGoods(ctx context.Context, in *GetAllGoodsReq, opts ...grpc.CallOption) (*GetAllGoodsRes, error) {
	out := new(GetAllGoodsRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetAllGoods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetAllGoodOptions(ctx context.Context, in *GetAllGoodOptionsReq, opts ...grpc.CallOption) (*GetAllGoodOptionsRes, error) {
	out := new(GetAllGoodOptionsRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetAllGoodOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddGoodClass(ctx context.Context, in *AddGoodClassReq, opts ...grpc.CallOption) (*AddGoodClassRes, error) {
	out := new(AddGoodClassRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddGoodClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddDeskClass(ctx context.Context, in *AddDeskClassReq, opts ...grpc.CallOption) (*AddDeskClassRes, error) {
	out := new(AddDeskClassRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddDeskClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddGood(ctx context.Context, in *AddGoodReq, opts ...grpc.CallOption) (*AddGoodRes, error) {
	out := new(AddGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) CancelGood(ctx context.Context, in *CancelGoodReq, opts ...grpc.CallOption) (*CancelGoodRes, error) {
	out := new(CancelGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/CancelGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddElement(ctx context.Context, in *AddElementReq, opts ...grpc.CallOption) (*AddElementRes, error) {
	out := new(AddElementRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddSpace(ctx context.Context, in *AddSpaceReq, opts ...grpc.CallOption) (*AddSpaceRes, error) {
	out := new(AddSpaceRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddDesk(ctx context.Context, in *AddDeskReq, opts ...grpc.CallOption) (*AddDeskRes, error) {
	out := new(AddDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) OrderGood(ctx context.Context, in *OrderGoodReq, opts ...grpc.CallOption) (*OrderGoodRes, error) {
	out := new(OrderGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/OrderGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) OrderDesk(ctx context.Context, in *OrderDeskReq, opts ...grpc.CallOption) (*OrderDeskRes, error) {
	out := new(OrderDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/OrderDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) ChangeDesk(ctx context.Context, in *ChangeDeskReq, opts ...grpc.CallOption) (*ChangeDeskRes, error) {
	out := new(ChangeDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/ChangeDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) CloseDesk(ctx context.Context, in *CloseDeskReq, opts ...grpc.CallOption) (*CloseDeskRes, error) {
	out := new(CloseDeskRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/CloseDesk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderRes, error) {
	out := new(GetOrderRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) CheckOut(ctx context.Context, in *CheckOutReq, opts ...grpc.CallOption) (*CheckOutRes, error) {
	out := new(CheckOutRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/CheckOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) AddFavorForGood(ctx context.Context, in *AddFavorForGoodReq, opts ...grpc.CallOption) (*AddFavorForGoodRes, error) {
	out := new(AddFavorForGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/AddFavorForGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) DeleteFavorForGood(ctx context.Context, in *DeleteFavorForGoodReq, opts ...grpc.CallOption) (*DeleteFavorForGoodRes, error) {
	out := new(DeleteFavorForGoodRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/DeleteFavorForGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mvpClient) GetAllDeskClasses(ctx context.Context, in *GetAllDeskClassesReq, opts ...grpc.CallOption) (*GetAllDeskClassesRes, error) {
	out := new(GetAllDeskClassesRes)
	err := c.cc.Invoke(ctx, "/pb.Mvp/GetAllDeskClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MvpServer is the server API for Mvp service.
// All implementations must embed UnimplementedMvpServer
// for forward compatibility
type MvpServer interface {
	GetAllGoodClasses(context.Context, *GetAllGoodClassesReq) (*GetAllGoodClassesRes, error)
	GetAllDesks(context.Context, *GetAllDesksReq) (*GetAllDesksRes, error)
	GetAllGoods(context.Context, *GetAllGoodsReq) (*GetAllGoodsRes, error)
	GetAllGoodOptions(context.Context, *GetAllGoodOptionsReq) (*GetAllGoodOptionsRes, error)
	AddGoodClass(context.Context, *AddGoodClassReq) (*AddGoodClassRes, error)
	AddDeskClass(context.Context, *AddDeskClassReq) (*AddDeskClassRes, error)
	AddGood(context.Context, *AddGoodReq) (*AddGoodRes, error)
	CancelGood(context.Context, *CancelGoodReq) (*CancelGoodRes, error)
	AddElement(context.Context, *AddElementReq) (*AddElementRes, error)
	AddSpace(context.Context, *AddSpaceReq) (*AddSpaceRes, error)
	AddDesk(context.Context, *AddDeskReq) (*AddDeskRes, error)
	OrderGood(context.Context, *OrderGoodReq) (*OrderGoodRes, error)
	OrderDesk(context.Context, *OrderDeskReq) (*OrderDeskRes, error)
	ChangeDesk(context.Context, *ChangeDeskReq) (*ChangeDeskRes, error)
	CloseDesk(context.Context, *CloseDeskReq) (*CloseDeskRes, error)
	GetOrder(context.Context, *GetOrderReq) (*GetOrderRes, error)
	CheckOut(context.Context, *CheckOutReq) (*CheckOutRes, error)
	AddFavorForGood(context.Context, *AddFavorForGoodReq) (*AddFavorForGoodRes, error)
	DeleteFavorForGood(context.Context, *DeleteFavorForGoodReq) (*DeleteFavorForGoodRes, error)
	GetAllDeskClasses(context.Context, *GetAllDeskClassesReq) (*GetAllDeskClassesRes, error)
	mustEmbedUnimplementedMvpServer()
}

// UnimplementedMvpServer must be embedded to have forward compatible implementations.
type UnimplementedMvpServer struct {
}

func (UnimplementedMvpServer) GetAllGoodClasses(context.Context, *GetAllGoodClassesReq) (*GetAllGoodClassesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGoodClasses not implemented")
}
func (UnimplementedMvpServer) GetAllDesks(context.Context, *GetAllDesksReq) (*GetAllDesksRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDesks not implemented")
}
func (UnimplementedMvpServer) GetAllGoods(context.Context, *GetAllGoodsReq) (*GetAllGoodsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGoods not implemented")
}
func (UnimplementedMvpServer) GetAllGoodOptions(context.Context, *GetAllGoodOptionsReq) (*GetAllGoodOptionsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGoodOptions not implemented")
}
func (UnimplementedMvpServer) AddGoodClass(context.Context, *AddGoodClassReq) (*AddGoodClassRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGoodClass not implemented")
}
func (UnimplementedMvpServer) AddDeskClass(context.Context, *AddDeskClassReq) (*AddDeskClassRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeskClass not implemented")
}
func (UnimplementedMvpServer) AddGood(context.Context, *AddGoodReq) (*AddGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGood not implemented")
}
func (UnimplementedMvpServer) CancelGood(context.Context, *CancelGoodReq) (*CancelGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelGood not implemented")
}
func (UnimplementedMvpServer) AddElement(context.Context, *AddElementReq) (*AddElementRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddElement not implemented")
}
func (UnimplementedMvpServer) AddSpace(context.Context, *AddSpaceReq) (*AddSpaceRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSpace not implemented")
}
func (UnimplementedMvpServer) AddDesk(context.Context, *AddDeskReq) (*AddDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDesk not implemented")
}
func (UnimplementedMvpServer) OrderGood(context.Context, *OrderGoodReq) (*OrderGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderGood not implemented")
}
func (UnimplementedMvpServer) OrderDesk(context.Context, *OrderDeskReq) (*OrderDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDesk not implemented")
}
func (UnimplementedMvpServer) ChangeDesk(context.Context, *ChangeDeskReq) (*ChangeDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeDesk not implemented")
}
func (UnimplementedMvpServer) CloseDesk(context.Context, *CloseDeskReq) (*CloseDeskRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloseDesk not implemented")
}
func (UnimplementedMvpServer) GetOrder(context.Context, *GetOrderReq) (*GetOrderRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedMvpServer) CheckOut(context.Context, *CheckOutReq) (*CheckOutRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckOut not implemented")
}
func (UnimplementedMvpServer) AddFavorForGood(context.Context, *AddFavorForGoodReq) (*AddFavorForGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFavorForGood not implemented")
}
func (UnimplementedMvpServer) DeleteFavorForGood(context.Context, *DeleteFavorForGoodReq) (*DeleteFavorForGoodRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFavorForGood not implemented")
}
func (UnimplementedMvpServer) GetAllDeskClasses(context.Context, *GetAllDeskClassesReq) (*GetAllDeskClassesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllDeskClasses not implemented")
}
func (UnimplementedMvpServer) mustEmbedUnimplementedMvpServer() {}

// UnsafeMvpServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MvpServer will
// result in compilation errors.
type UnsafeMvpServer interface {
	mustEmbedUnimplementedMvpServer()
}

func RegisterMvpServer(s grpc.ServiceRegistrar, srv MvpServer) {
	s.RegisterService(&_Mvp_serviceDesc, srv)
}

func _Mvp_GetAllGoodClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGoodClassesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetAllGoodClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetAllGoodClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetAllGoodClasses(ctx, req.(*GetAllGoodClassesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetAllDesks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDesksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetAllDesks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetAllDesks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetAllDesks(ctx, req.(*GetAllDesksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetAllGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGoodsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetAllGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetAllGoods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetAllGoods(ctx, req.(*GetAllGoodsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetAllGoodOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGoodOptionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetAllGoodOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetAllGoodOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetAllGoodOptions(ctx, req.(*GetAllGoodOptionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddGoodClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodClassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddGoodClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddGoodClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddGoodClass(ctx, req.(*AddGoodClassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddDeskClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeskClassReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddDeskClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddDeskClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddDeskClass(ctx, req.(*AddDeskClassReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddGood(ctx, req.(*AddGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_CancelGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).CancelGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/CancelGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).CancelGood(ctx, req.(*CancelGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddElementReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddElement(ctx, req.(*AddElementReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSpaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddSpace(ctx, req.(*AddSpaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddDesk(ctx, req.(*AddDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_OrderGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).OrderGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/OrderGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).OrderGood(ctx, req.(*OrderGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_OrderDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).OrderDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/OrderDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).OrderDesk(ctx, req.(*OrderDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_ChangeDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).ChangeDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/ChangeDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).ChangeDesk(ctx, req.(*ChangeDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_CloseDesk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseDeskReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).CloseDesk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/CloseDesk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).CloseDesk(ctx, req.(*CloseDeskReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetOrder(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_CheckOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckOutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).CheckOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/CheckOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).CheckOut(ctx, req.(*CheckOutReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_AddFavorForGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFavorForGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).AddFavorForGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/AddFavorForGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).AddFavorForGood(ctx, req.(*AddFavorForGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_DeleteFavorForGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFavorForGoodReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).DeleteFavorForGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/DeleteFavorForGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).DeleteFavorForGood(ctx, req.(*DeleteFavorForGoodReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mvp_GetAllDeskClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDeskClassesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MvpServer).GetAllDeskClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Mvp/GetAllDeskClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MvpServer).GetAllDeskClasses(ctx, req.(*GetAllDeskClassesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mvp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Mvp",
	HandlerType: (*MvpServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllGoodClasses",
			Handler:    _Mvp_GetAllGoodClasses_Handler,
		},
		{
			MethodName: "GetAllDesks",
			Handler:    _Mvp_GetAllDesks_Handler,
		},
		{
			MethodName: "GetAllGoods",
			Handler:    _Mvp_GetAllGoods_Handler,
		},
		{
			MethodName: "GetAllGoodOptions",
			Handler:    _Mvp_GetAllGoodOptions_Handler,
		},
		{
			MethodName: "AddGoodClass",
			Handler:    _Mvp_AddGoodClass_Handler,
		},
		{
			MethodName: "AddDeskClass",
			Handler:    _Mvp_AddDeskClass_Handler,
		},
		{
			MethodName: "AddGood",
			Handler:    _Mvp_AddGood_Handler,
		},
		{
			MethodName: "CancelGood",
			Handler:    _Mvp_CancelGood_Handler,
		},
		{
			MethodName: "AddElement",
			Handler:    _Mvp_AddElement_Handler,
		},
		{
			MethodName: "AddSpace",
			Handler:    _Mvp_AddSpace_Handler,
		},
		{
			MethodName: "AddDesk",
			Handler:    _Mvp_AddDesk_Handler,
		},
		{
			MethodName: "OrderGood",
			Handler:    _Mvp_OrderGood_Handler,
		},
		{
			MethodName: "OrderDesk",
			Handler:    _Mvp_OrderDesk_Handler,
		},
		{
			MethodName: "ChangeDesk",
			Handler:    _Mvp_ChangeDesk_Handler,
		},
		{
			MethodName: "CloseDesk",
			Handler:    _Mvp_CloseDesk_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _Mvp_GetOrder_Handler,
		},
		{
			MethodName: "CheckOut",
			Handler:    _Mvp_CheckOut_Handler,
		},
		{
			MethodName: "AddFavorForGood",
			Handler:    _Mvp_AddFavorForGood_Handler,
		},
		{
			MethodName: "DeleteFavorForGood",
			Handler:    _Mvp_DeleteFavorForGood_Handler,
		},
		{
			MethodName: "GetAllDeskClasses",
			Handler:    _Mvp_GetAllDeskClasses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/pb/mvp.proto",
}
