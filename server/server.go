package server

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))

	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *HandlerBaseOnMap
}

func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handlerFunc(ctx)
	//})
	key := s.handler.key(method, pattern)

	// TODO 解决重复注册的问题
	s.handler.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Start(address string) error {

	http.Handle("/", s.handler)

	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}
}

func Signup(ctx *Context) {

	req := &signupReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{Data: 123}

	err = ctx.OkJson(resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err) // 写入失败了就不必重试了，这里应该写日志
		return
	}
}

type signupReq struct {
	Email string `json:"email"`
}
type commonResponse struct {
	Data int `json:"data"`
}
