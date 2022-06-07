// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: core/content.proto

package core

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentServiceClient interface {
	// コンテンツリストのリストを取得します。
	// 取得するコンテンツリストのチームIDを指定した GetContentListsRequest を渡します。
	// 指定されたチームIDのコンテンツリストのリストが設定された GetContentListsResponse が返ります。
	GetContentLists(ctx context.Context, in *GetContentListsRequest, opts ...grpc.CallOption) (*GetContentListsResponse, error)
	// コンテンツリストを取得します。
	// 取得するコンテンツリストのコンテンツリストIDを指定した GetContentListRequest を渡します。
	// コンテンツリストが存在する場合、コンテンツリストが設定された GetContentListResponse が返ります。
	GetContentList(ctx context.Context, in *GetContentListRequest, opts ...grpc.CallOption) (*GetContentListResponse, error)
	// 新しくコンテンツリストを作成します。
	// コンテンツリスト名とチームIDを指定した CreateContentListRequest を渡します。
	// コンテンツリストの作成に成功すると、作成されたコンテンツリストが設定された CreateContentListResponse が返ります。
	CreateContentList(ctx context.Context, in *CreateContentListRequest, opts ...grpc.CallOption) (*CreateContentListResponse, error)
	// コンテンツリストを更新します。
	// 更新するコンテンツリストのコンテンツリストIDと、新しいコンテンツリスト名を指定した ContentListUpdateRequest を渡します。
	// コンテンツリストの作成に成功すると、ContentList が返ります。
	UpdateContentList(ctx context.Context, in *UpdateContentListRequest, opts ...grpc.CallOption) (*UpdateContentListResponse, error)
	// コンテンツリストを削除します。
	// 削除するコンテンツリストのコンテンツリストIDを指定した ContentListRequest を渡します。
	DeleteContentList(ctx context.Context, in *DeleteContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// コンテンツをコンテンツリストに追加します。
	// チームIDと追加するコンテンツのコンテンツID、追加先のコンテンツリストのコンテンツリストIDを指定した AddContentToContentListRequest を渡します。
	AddContentToContentList(ctx context.Context, in *AddContentToContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// コンテンツをコンテンツリストから削除します。
	// チームIDと削除するコンテンツのコンテンツID、削除元のコンテンツリストのコンテンツリストIDを指定した RemoveContentFromContentListRequest を渡します。
	RemoveContentFromContentList(ctx context.Context, in *RemoveContentFromContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// コンテンツリスト内のコンテンツを移動します。
	// チームIDと移動するコンテンツのコンテンツID、移動するコンテンツリストのコンテンツリストID、移動先のインデックスを指定した MoveContentInContentListRequest を渡します。
	MoveContentInContentList(ctx context.Context, in *MoveContentInContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// 全てのコンテンツのリストを取得します。
	// 取得するコンテンツのチームIDを指定した GetAllContentsRequest を渡します。
	// contents_types に1つ以上のコンテンツタイプを指定すると、指定されたコンテンツタイプのコンテンツのみ取得します。
	// 取得したコンテンツのリストが設定された GetAllContentsResponse が返ります。
	GetAllContents(ctx context.Context, in *GetAllContentsRequest, opts ...grpc.CallOption) (*GetAllContentsResponse, error)
	// 指定されたコンテンツリストに追加されているコンテンツのリストを取得します。
	// 取得するコンテンツのチームIDとコンテンツリストIDを指定した GetContentsRequest を渡します。
	// contents_types に1つ以上のコンテンツタイプを指定すると、指定されたコンテンツタイプのコンテンツのみ取得します。
	// 取得したコンテンツのリストが設定された GetContentsResponse が返ります。
	GetContents(ctx context.Context, in *GetContentsRequest, opts ...grpc.CallOption) (*GetContentsResponse, error)
	// コンテンツを取得します。
	// 取得するコンテンツのコンテンツIDを指定した GetContentRequest を渡します。
	// コンテンツが存在する場合、コンテンツが設定された GetContentResponse が返ります。
	GetContent(ctx context.Context, in *GetContentRequest, opts ...grpc.CallOption) (*GetContentResponse, error)
	// コンテンツをアップロードします。
	// チームID、アップロードするコンテンツのコンテンツ名とコンテンツタイプを指定した UploadContentRequest を渡します。
	// アップロードを行うための URL が設定された ContentUploadURL が返ります。
	// このURLにコンテンツデータをHTTP POSTメソッドで転送することで、アップロードが行われます。
	// アップロード完了後、FinishUpload メソッドでアップロードの完了を登録します。
	UploadContent(ctx context.Context, in *UploadContentRequest, opts ...grpc.CallOption) (*UploadContentResponse, error)
	// コンテンツのアップロード完了を登録します。
	// チームID、コンテンツID、アップロードしたコンテンツのMD5ハッシュを指定した FinishUploadRequest を渡します。
	// 登録に成功すると Content が返ります。
	FinishUploadContent(ctx context.Context, in *FinishUploadContentRequest, opts ...grpc.CallOption) (*FinishUploadContentResponse, error)
	// コンテンツをダウンロードするためのを取得します。
	// ダウンロードするコンテンツのチームID、コンテンツIDを指定した、DownloadContentRequest を渡します。
	// コンテンツが存在する場合、コンテンツとダウンロードURLが設定された DownloadContentResponse が返ります。
	DownloadContent(ctx context.Context, in *DownloadContentRequest, opts ...grpc.CallOption) (*DownloadContentResponse, error)
	// コンテンツを更新します。
	// 更新するコンテンツのコンテンツIDと、新しいコンテンツ名を指定した ContentUpdateRequest を渡します。
	// コンテンツの作成に成功すると、Content が返ります。
	UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...grpc.CallOption) (*UpdateContentResponse, error)
	// コンテンツを削除します。
	// 削除するコンテンツのコンテンツIDを指定した ContentRequest を渡します。
	DeleteContent(ctx context.Context, in *DeleteContentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) GetContentLists(ctx context.Context, in *GetContentListsRequest, opts ...grpc.CallOption) (*GetContentListsResponse, error) {
	out := new(GetContentListsResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/GetContentLists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContentList(ctx context.Context, in *GetContentListRequest, opts ...grpc.CallOption) (*GetContentListResponse, error) {
	out := new(GetContentListResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/GetContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) CreateContentList(ctx context.Context, in *CreateContentListRequest, opts ...grpc.CallOption) (*CreateContentListResponse, error) {
	out := new(CreateContentListResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/CreateContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateContentList(ctx context.Context, in *UpdateContentListRequest, opts ...grpc.CallOption) (*UpdateContentListResponse, error) {
	out := new(UpdateContentListResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/UpdateContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteContentList(ctx context.Context, in *DeleteContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/DeleteContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) AddContentToContentList(ctx context.Context, in *AddContentToContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/AddContentToContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) RemoveContentFromContentList(ctx context.Context, in *RemoveContentFromContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/RemoveContentFromContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) MoveContentInContentList(ctx context.Context, in *MoveContentInContentListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/MoveContentInContentList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetAllContents(ctx context.Context, in *GetAllContentsRequest, opts ...grpc.CallOption) (*GetAllContentsResponse, error) {
	out := new(GetAllContentsResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/GetAllContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContents(ctx context.Context, in *GetContentsRequest, opts ...grpc.CallOption) (*GetContentsResponse, error) {
	out := new(GetContentsResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/GetContents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetContent(ctx context.Context, in *GetContentRequest, opts ...grpc.CallOption) (*GetContentResponse, error) {
	out := new(GetContentResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UploadContent(ctx context.Context, in *UploadContentRequest, opts ...grpc.CallOption) (*UploadContentResponse, error) {
	out := new(UploadContentResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/UploadContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) FinishUploadContent(ctx context.Context, in *FinishUploadContentRequest, opts ...grpc.CallOption) (*FinishUploadContentResponse, error) {
	out := new(FinishUploadContentResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/FinishUploadContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DownloadContent(ctx context.Context, in *DownloadContentRequest, opts ...grpc.CallOption) (*DownloadContentResponse, error) {
	out := new(DownloadContentResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/DownloadContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) UpdateContent(ctx context.Context, in *UpdateContentRequest, opts ...grpc.CallOption) (*UpdateContentResponse, error) {
	out := new(UpdateContentResponse)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/UpdateContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) DeleteContent(ctx context.Context, in *DeleteContentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/at.core.ContentService/DeleteContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
// All implementations must embed UnimplementedContentServiceServer
// for forward compatibility
type ContentServiceServer interface {
	// コンテンツリストのリストを取得します。
	// 取得するコンテンツリストのチームIDを指定した GetContentListsRequest を渡します。
	// 指定されたチームIDのコンテンツリストのリストが設定された GetContentListsResponse が返ります。
	GetContentLists(context.Context, *GetContentListsRequest) (*GetContentListsResponse, error)
	// コンテンツリストを取得します。
	// 取得するコンテンツリストのコンテンツリストIDを指定した GetContentListRequest を渡します。
	// コンテンツリストが存在する場合、コンテンツリストが設定された GetContentListResponse が返ります。
	GetContentList(context.Context, *GetContentListRequest) (*GetContentListResponse, error)
	// 新しくコンテンツリストを作成します。
	// コンテンツリスト名とチームIDを指定した CreateContentListRequest を渡します。
	// コンテンツリストの作成に成功すると、作成されたコンテンツリストが設定された CreateContentListResponse が返ります。
	CreateContentList(context.Context, *CreateContentListRequest) (*CreateContentListResponse, error)
	// コンテンツリストを更新します。
	// 更新するコンテンツリストのコンテンツリストIDと、新しいコンテンツリスト名を指定した ContentListUpdateRequest を渡します。
	// コンテンツリストの作成に成功すると、ContentList が返ります。
	UpdateContentList(context.Context, *UpdateContentListRequest) (*UpdateContentListResponse, error)
	// コンテンツリストを削除します。
	// 削除するコンテンツリストのコンテンツリストIDを指定した ContentListRequest を渡します。
	DeleteContentList(context.Context, *DeleteContentListRequest) (*emptypb.Empty, error)
	// コンテンツをコンテンツリストに追加します。
	// チームIDと追加するコンテンツのコンテンツID、追加先のコンテンツリストのコンテンツリストIDを指定した AddContentToContentListRequest を渡します。
	AddContentToContentList(context.Context, *AddContentToContentListRequest) (*emptypb.Empty, error)
	// コンテンツをコンテンツリストから削除します。
	// チームIDと削除するコンテンツのコンテンツID、削除元のコンテンツリストのコンテンツリストIDを指定した RemoveContentFromContentListRequest を渡します。
	RemoveContentFromContentList(context.Context, *RemoveContentFromContentListRequest) (*emptypb.Empty, error)
	// コンテンツリスト内のコンテンツを移動します。
	// チームIDと移動するコンテンツのコンテンツID、移動するコンテンツリストのコンテンツリストID、移動先のインデックスを指定した MoveContentInContentListRequest を渡します。
	MoveContentInContentList(context.Context, *MoveContentInContentListRequest) (*emptypb.Empty, error)
	// 全てのコンテンツのリストを取得します。
	// 取得するコンテンツのチームIDを指定した GetAllContentsRequest を渡します。
	// contents_types に1つ以上のコンテンツタイプを指定すると、指定されたコンテンツタイプのコンテンツのみ取得します。
	// 取得したコンテンツのリストが設定された GetAllContentsResponse が返ります。
	GetAllContents(context.Context, *GetAllContentsRequest) (*GetAllContentsResponse, error)
	// 指定されたコンテンツリストに追加されているコンテンツのリストを取得します。
	// 取得するコンテンツのチームIDとコンテンツリストIDを指定した GetContentsRequest を渡します。
	// contents_types に1つ以上のコンテンツタイプを指定すると、指定されたコンテンツタイプのコンテンツのみ取得します。
	// 取得したコンテンツのリストが設定された GetContentsResponse が返ります。
	GetContents(context.Context, *GetContentsRequest) (*GetContentsResponse, error)
	// コンテンツを取得します。
	// 取得するコンテンツのコンテンツIDを指定した GetContentRequest を渡します。
	// コンテンツが存在する場合、コンテンツが設定された GetContentResponse が返ります。
	GetContent(context.Context, *GetContentRequest) (*GetContentResponse, error)
	// コンテンツをアップロードします。
	// チームID、アップロードするコンテンツのコンテンツ名とコンテンツタイプを指定した UploadContentRequest を渡します。
	// アップロードを行うための URL が設定された ContentUploadURL が返ります。
	// このURLにコンテンツデータをHTTP POSTメソッドで転送することで、アップロードが行われます。
	// アップロード完了後、FinishUpload メソッドでアップロードの完了を登録します。
	UploadContent(context.Context, *UploadContentRequest) (*UploadContentResponse, error)
	// コンテンツのアップロード完了を登録します。
	// チームID、コンテンツID、アップロードしたコンテンツのMD5ハッシュを指定した FinishUploadRequest を渡します。
	// 登録に成功すると Content が返ります。
	FinishUploadContent(context.Context, *FinishUploadContentRequest) (*FinishUploadContentResponse, error)
	// コンテンツをダウンロードするためのを取得します。
	// ダウンロードするコンテンツのチームID、コンテンツIDを指定した、DownloadContentRequest を渡します。
	// コンテンツが存在する場合、コンテンツとダウンロードURLが設定された DownloadContentResponse が返ります。
	DownloadContent(context.Context, *DownloadContentRequest) (*DownloadContentResponse, error)
	// コンテンツを更新します。
	// 更新するコンテンツのコンテンツIDと、新しいコンテンツ名を指定した ContentUpdateRequest を渡します。
	// コンテンツの作成に成功すると、Content が返ります。
	UpdateContent(context.Context, *UpdateContentRequest) (*UpdateContentResponse, error)
	// コンテンツを削除します。
	// 削除するコンテンツのコンテンツIDを指定した ContentRequest を渡します。
	DeleteContent(context.Context, *DeleteContentRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedContentServiceServer()
}

// UnimplementedContentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentServiceServer struct {
}

func (UnimplementedContentServiceServer) GetContentLists(context.Context, *GetContentListsRequest) (*GetContentListsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentLists not implemented")
}
func (UnimplementedContentServiceServer) GetContentList(context.Context, *GetContentListRequest) (*GetContentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContentList not implemented")
}
func (UnimplementedContentServiceServer) CreateContentList(context.Context, *CreateContentListRequest) (*CreateContentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContentList not implemented")
}
func (UnimplementedContentServiceServer) UpdateContentList(context.Context, *UpdateContentListRequest) (*UpdateContentListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContentList not implemented")
}
func (UnimplementedContentServiceServer) DeleteContentList(context.Context, *DeleteContentListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContentList not implemented")
}
func (UnimplementedContentServiceServer) AddContentToContentList(context.Context, *AddContentToContentListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddContentToContentList not implemented")
}
func (UnimplementedContentServiceServer) RemoveContentFromContentList(context.Context, *RemoveContentFromContentListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveContentFromContentList not implemented")
}
func (UnimplementedContentServiceServer) MoveContentInContentList(context.Context, *MoveContentInContentListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveContentInContentList not implemented")
}
func (UnimplementedContentServiceServer) GetAllContents(context.Context, *GetAllContentsRequest) (*GetAllContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllContents not implemented")
}
func (UnimplementedContentServiceServer) GetContents(context.Context, *GetContentsRequest) (*GetContentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContents not implemented")
}
func (UnimplementedContentServiceServer) GetContent(context.Context, *GetContentRequest) (*GetContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedContentServiceServer) UploadContent(context.Context, *UploadContentRequest) (*UploadContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadContent not implemented")
}
func (UnimplementedContentServiceServer) FinishUploadContent(context.Context, *FinishUploadContentRequest) (*FinishUploadContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishUploadContent not implemented")
}
func (UnimplementedContentServiceServer) DownloadContent(context.Context, *DownloadContentRequest) (*DownloadContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadContent not implemented")
}
func (UnimplementedContentServiceServer) UpdateContent(context.Context, *UpdateContentRequest) (*UpdateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContent not implemented")
}
func (UnimplementedContentServiceServer) DeleteContent(context.Context, *DeleteContentRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}

// UnsafeContentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentServiceServer will
// result in compilation errors.
type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_GetContentLists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentListsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContentLists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/GetContentLists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContentLists(ctx, req.(*GetContentListsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/GetContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContentList(ctx, req.(*GetContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_CreateContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).CreateContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/CreateContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).CreateContentList(ctx, req.(*CreateContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UpdateContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UpdateContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/UpdateContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UpdateContentList(ctx, req.(*UpdateContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/DeleteContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteContentList(ctx, req.(*DeleteContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_AddContentToContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddContentToContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).AddContentToContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/AddContentToContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).AddContentToContentList(ctx, req.(*AddContentToContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_RemoveContentFromContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveContentFromContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).RemoveContentFromContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/RemoveContentFromContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).RemoveContentFromContentList(ctx, req.(*RemoveContentFromContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_MoveContentInContentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveContentInContentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).MoveContentInContentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/MoveContentInContentList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).MoveContentInContentList(ctx, req.(*MoveContentInContentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetAllContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetAllContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/GetAllContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetAllContents(ctx, req.(*GetAllContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetContents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/GetContents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContents(ctx, req.(*GetContentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetContent(ctx, req.(*GetContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UploadContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UploadContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/UploadContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UploadContent(ctx, req.(*UploadContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_FinishUploadContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishUploadContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).FinishUploadContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/FinishUploadContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).FinishUploadContent(ctx, req.(*FinishUploadContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DownloadContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DownloadContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/DownloadContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DownloadContent(ctx, req.(*DownloadContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_UpdateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).UpdateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/UpdateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).UpdateContent(ctx, req.(*UpdateContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/at.core.ContentService/DeleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).DeleteContent(ctx, req.(*DeleteContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentService_ServiceDesc is the grpc.ServiceDesc for ContentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "at.core.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContentLists",
			Handler:    _ContentService_GetContentLists_Handler,
		},
		{
			MethodName: "GetContentList",
			Handler:    _ContentService_GetContentList_Handler,
		},
		{
			MethodName: "CreateContentList",
			Handler:    _ContentService_CreateContentList_Handler,
		},
		{
			MethodName: "UpdateContentList",
			Handler:    _ContentService_UpdateContentList_Handler,
		},
		{
			MethodName: "DeleteContentList",
			Handler:    _ContentService_DeleteContentList_Handler,
		},
		{
			MethodName: "AddContentToContentList",
			Handler:    _ContentService_AddContentToContentList_Handler,
		},
		{
			MethodName: "RemoveContentFromContentList",
			Handler:    _ContentService_RemoveContentFromContentList_Handler,
		},
		{
			MethodName: "MoveContentInContentList",
			Handler:    _ContentService_MoveContentInContentList_Handler,
		},
		{
			MethodName: "GetAllContents",
			Handler:    _ContentService_GetAllContents_Handler,
		},
		{
			MethodName: "GetContents",
			Handler:    _ContentService_GetContents_Handler,
		},
		{
			MethodName: "GetContent",
			Handler:    _ContentService_GetContent_Handler,
		},
		{
			MethodName: "UploadContent",
			Handler:    _ContentService_UploadContent_Handler,
		},
		{
			MethodName: "FinishUploadContent",
			Handler:    _ContentService_FinishUploadContent_Handler,
		},
		{
			MethodName: "DownloadContent",
			Handler:    _ContentService_DownloadContent_Handler,
		},
		{
			MethodName: "UpdateContent",
			Handler:    _ContentService_UpdateContent_Handler,
		},
		{
			MethodName: "DeleteContent",
			Handler:    _ContentService_DeleteContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/content.proto",
}
