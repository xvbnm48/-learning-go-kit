package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/xvbnm48/-learning-go-kit/service"
)

type Endpoints struct {
	Add        endpoint.Endpoint
	Subtract   endpoint.Endpoint
	Multiply   endpoint.Endpoint
	Divide     endpoint.Endpoint
	Palindrome endpoint.Endpoint
}

type PalindromeReq struct {
	Word string
}

type PalindromeResp struct {
	Result bool
}

type MathReq struct {
	NumA float32
	NumB float32
}

type MathResp struct {
	Result float32
}

func MakeEndpoints(s service.Service) Endpoints {
	return Endpoints{
		Add:        makeAddEndpoint(s),
		Subtract:   makeSubtract(s),
		Multiply:   makeMultiply(s),
		Divide:     makeDevide(s),
		Palindrome: makePalindrome(s),
	}
}

func makePalindrome(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PalindromeReq)
		result, _ := s.Palindrome(ctx, req.Word)
		return PalindromeResp{Result: result}, nil
	}
}

func makeAddEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeSubtract(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Subtract(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeMultiply(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Multiply(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeDevide(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Divide(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}
