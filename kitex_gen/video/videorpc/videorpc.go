// Code generated by Kitex v0.5.1. DO NOT EDIT.

package videorpc

import (
	"context"
	"fmt"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	video "tiktok_micro/kitex_gen/video"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoRpcServiceInfo
}

var videoRpcServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoRpc"
	handlerType := (*video.VideoRpc)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetPublishList":  kitex.NewMethodInfo(getPublishListHandler, newGetPublishListArgs, newGetPublishListResult, false),
		"GetFeed":         kitex.NewMethodInfo(getFeedHandler, newGetFeedArgs, newGetFeedResult, false),
		"CommentAction":   kitex.NewMethodInfo(commentActionHandler, newCommentActionArgs, newCommentActionResult, false),
		"GetCommentList":  kitex.NewMethodInfo(getCommentListHandler, newGetCommentListArgs, newGetCommentListResult, false),
		"FavoriteAction":  kitex.NewMethodInfo(favoriteActionHandler, newFavoriteActionArgs, newFavoriteActionResult, false),
		"GetFavoriteList": kitex.NewMethodInfo(getFavoriteListHandler, newGetFavoriteListArgs, newGetFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func getPublishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PublishListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoRpc).GetPublishList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPublishListArgs:
		success, err := handler.(video.VideoRpc).GetPublishList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPublishListResult)
		realResult.Success = success
	}
	return nil
}
func newGetPublishListArgs() interface{} {
	return &GetPublishListArgs{}
}

func newGetPublishListResult() interface{} {
	return &GetPublishListResult{}
}

type GetPublishListArgs struct {
	Req *video.PublishListReq
}

func (p *GetPublishListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.PublishListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetPublishListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetPublishListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetPublishListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPublishListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPublishListArgs) Unmarshal(in []byte) error {
	msg := new(video.PublishListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPublishListArgs_Req_DEFAULT *video.PublishListReq

func (p *GetPublishListArgs) GetReq() *video.PublishListReq {
	if !p.IsSetReq() {
		return GetPublishListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPublishListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetPublishListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetPublishListResult struct {
	Success *video.PublishListResp
}

var GetPublishListResult_Success_DEFAULT *video.PublishListResp

func (p *GetPublishListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.PublishListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetPublishListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetPublishListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetPublishListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPublishListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPublishListResult) Unmarshal(in []byte) error {
	msg := new(video.PublishListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPublishListResult) GetSuccess() *video.PublishListResp {
	if !p.IsSetSuccess() {
		return GetPublishListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPublishListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PublishListResp)
}

func (p *GetPublishListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetPublishListResult) GetResult() interface{} {
	return p.Success
}

func getFeedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FeedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoRpc).GetFeed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFeedArgs:
		success, err := handler.(video.VideoRpc).GetFeed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFeedResult)
		realResult.Success = success
	}
	return nil
}
func newGetFeedArgs() interface{} {
	return &GetFeedArgs{}
}

func newGetFeedResult() interface{} {
	return &GetFeedResult{}
}

type GetFeedArgs struct {
	Req *video.FeedReq
}

func (p *GetFeedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.FeedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFeedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFeedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFeedArgs) Unmarshal(in []byte) error {
	msg := new(video.FeedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFeedArgs_Req_DEFAULT *video.FeedReq

func (p *GetFeedArgs) GetReq() *video.FeedReq {
	if !p.IsSetReq() {
		return GetFeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFeedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFeedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFeedResult struct {
	Success *video.FeedResp
}

var GetFeedResult_Success_DEFAULT *video.FeedResp

func (p *GetFeedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.FeedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFeedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFeedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFeedResult) Unmarshal(in []byte) error {
	msg := new(video.FeedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFeedResult) GetSuccess() *video.FeedResp {
	if !p.IsSetSuccess() {
		return GetFeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FeedResp)
}

func (p *GetFeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFeedResult) GetResult() interface{} {
	return p.Success
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.CommentReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoRpc).CommentAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CommentActionArgs:
		success, err := handler.(video.VideoRpc).CommentAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CommentActionResult)
		realResult.Success = success
	}
	return nil
}
func newCommentActionArgs() interface{} {
	return &CommentActionArgs{}
}

func newCommentActionResult() interface{} {
	return &CommentActionResult{}
}

type CommentActionArgs struct {
	Req *video.CommentReq
}

func (p *CommentActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.CommentReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CommentActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CommentActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CommentActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CommentActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CommentActionArgs) Unmarshal(in []byte) error {
	msg := new(video.CommentReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CommentActionArgs_Req_DEFAULT *video.CommentReq

func (p *CommentActionArgs) GetReq() *video.CommentReq {
	if !p.IsSetReq() {
		return CommentActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CommentActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CommentActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CommentActionResult struct {
	Success *video.CommentResp
}

var CommentActionResult_Success_DEFAULT *video.CommentResp

func (p *CommentActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.CommentResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CommentActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CommentActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CommentActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CommentActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CommentActionResult) Unmarshal(in []byte) error {
	msg := new(video.CommentResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CommentActionResult) GetSuccess() *video.CommentResp {
	if !p.IsSetSuccess() {
		return CommentActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CommentActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.CommentResp)
}

func (p *CommentActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CommentActionResult) GetResult() interface{} {
	return p.Success
}

func getCommentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.CommentListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoRpc).GetCommentList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentListArgs:
		success, err := handler.(video.VideoRpc).GetCommentList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentListResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentListArgs() interface{} {
	return &GetCommentListArgs{}
}

func newGetCommentListResult() interface{} {
	return &GetCommentListResult{}
}

type GetCommentListArgs struct {
	Req *video.CommentListReq
}

func (p *GetCommentListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.CommentListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetCommentListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetCommentListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetCommentListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentListArgs) Unmarshal(in []byte) error {
	msg := new(video.CommentListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentListArgs_Req_DEFAULT *video.CommentListReq

func (p *GetCommentListArgs) GetReq() *video.CommentListReq {
	if !p.IsSetReq() {
		return GetCommentListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetCommentListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetCommentListResult struct {
	Success *video.CommentListResp
}

var GetCommentListResult_Success_DEFAULT *video.CommentListResp

func (p *GetCommentListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.CommentListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetCommentListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetCommentListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetCommentListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentListResult) Unmarshal(in []byte) error {
	msg := new(video.CommentListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentListResult) GetSuccess() *video.CommentListResp {
	if !p.IsSetSuccess() {
		return GetCommentListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.CommentListResp)
}

func (p *GetCommentListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetCommentListResult) GetResult() interface{} {
	return p.Success
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FavoriteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoRpc).FavoriteAction(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteActionArgs:
		success, err := handler.(video.VideoRpc).FavoriteAction(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteActionResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteActionArgs() interface{} {
	return &FavoriteActionArgs{}
}

func newFavoriteActionResult() interface{} {
	return &FavoriteActionResult{}
}

type FavoriteActionArgs struct {
	Req *video.FavoriteReq
}

func (p *FavoriteActionArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.FavoriteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *FavoriteActionArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *FavoriteActionArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *FavoriteActionArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteActionArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteActionArgs) Unmarshal(in []byte) error {
	msg := new(video.FavoriteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteActionArgs_Req_DEFAULT *video.FavoriteReq

func (p *FavoriteActionArgs) GetReq() *video.FavoriteReq {
	if !p.IsSetReq() {
		return FavoriteActionArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteActionArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *FavoriteActionArgs) GetFirstArgument() interface{} {
	return p.Req
}

type FavoriteActionResult struct {
	Success *video.FavoriteResp
}

var FavoriteActionResult_Success_DEFAULT *video.FavoriteResp

func (p *FavoriteActionResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.FavoriteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *FavoriteActionResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *FavoriteActionResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *FavoriteActionResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteActionResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteActionResult) Unmarshal(in []byte) error {
	msg := new(video.FavoriteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteActionResult) GetSuccess() *video.FavoriteResp {
	if !p.IsSetSuccess() {
		return FavoriteActionResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteActionResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FavoriteResp)
}

func (p *FavoriteActionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FavoriteActionResult) GetResult() interface{} {
	return p.Success
}

func getFavoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FavoriteListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoRpc).GetFavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteListArgs:
		success, err := handler.(video.VideoRpc).GetFavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteListArgs() interface{} {
	return &GetFavoriteListArgs{}
}

func newGetFavoriteListResult() interface{} {
	return &GetFavoriteListResult{}
}

type GetFavoriteListArgs struct {
	Req *video.FavoriteListReq
}

func (p *GetFavoriteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(video.FavoriteListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetFavoriteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetFavoriteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetFavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(video.FavoriteListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteListArgs_Req_DEFAULT *video.FavoriteListReq

func (p *GetFavoriteListArgs) GetReq() *video.FavoriteListReq {
	if !p.IsSetReq() {
		return GetFavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetFavoriteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetFavoriteListResult struct {
	Success *video.FavoriteListResp
}

var GetFavoriteListResult_Success_DEFAULT *video.FavoriteListResp

func (p *GetFavoriteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(video.FavoriteListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetFavoriteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetFavoriteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetFavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteListResult) Unmarshal(in []byte) error {
	msg := new(video.FavoriteListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteListResult) GetSuccess() *video.FavoriteListResp {
	if !p.IsSetSuccess() {
		return GetFavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FavoriteListResp)
}

func (p *GetFavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetFavoriteListResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetPublishList(ctx context.Context, Req *video.PublishListReq) (r *video.PublishListResp, err error) {
	var _args GetPublishListArgs
	_args.Req = Req
	var _result GetPublishListResult
	if err = p.c.Call(ctx, "GetPublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFeed(ctx context.Context, Req *video.FeedReq) (r *video.FeedResp, err error) {
	var _args GetFeedArgs
	_args.Req = Req
	var _result GetFeedResult
	if err = p.c.Call(ctx, "GetFeed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentAction(ctx context.Context, Req *video.CommentReq) (r *video.CommentResp, err error) {
	var _args CommentActionArgs
	_args.Req = Req
	var _result CommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetCommentList(ctx context.Context, Req *video.CommentListReq) (r *video.CommentListResp, err error) {
	var _args GetCommentListArgs
	_args.Req = Req
	var _result GetCommentListResult
	if err = p.c.Call(ctx, "GetCommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteAction(ctx context.Context, Req *video.FavoriteReq) (r *video.FavoriteResp, err error) {
	var _args FavoriteActionArgs
	_args.Req = Req
	var _result FavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteList(ctx context.Context, Req *video.FavoriteListReq) (r *video.FavoriteListResp, err error) {
	var _args GetFavoriteListArgs
	_args.Req = Req
	var _result GetFavoriteListResult
	if err = p.c.Call(ctx, "GetFavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
