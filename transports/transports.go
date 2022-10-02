package transports

import (
	"context"

	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/xvbnm48/-learning-go-kit/endpoints"
	"github.com/xvbnm48/-learning-go-kit/pb"
)

type gRPCServer struct {
	add        gt.Handler
	substract  gt.Handler
	multiply   gt.Handler
	divide     gt.Handler
	palindrome gt.Handler
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: gt.NewServer(
			endpoints.Add,
			decodeMathRequest,
			encodeMathResponse,
		),
		substract: gt.NewServer(
			endpoints.Subtract,
			decodeMathRequest,
			encodeMathResponse,
		),
		multiply: gt.NewServer(
			endpoints.Multiply,
			decodeMathRequest,
			encodeMathResponse,
		),
		divide: gt.NewServer(
			endpoints.Divide,
			decodeMathRequest,
			encodeMathResponse,
		),
		palindrome: gt.NewServer(
			endpoints.Palindrome,
			decodePalindromeRequest,
			encodePalindromeResponse,
		),
	}
}

func (s *gRPCServer) Palindrome(ctx context.Context, req *pb.PalindromeRequest) (*pb.PalindromeResponse, error) {
	_, resp, err := s.palindrome.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.PalindromeResponse), nil
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Subtract(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.substract.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Multiply(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.multiply.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Divide(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.divide.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoints.MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.MathResponse{Result: resp.Result}, nil
}

func decodePalindromeRequest(_ context.Context, reqest interface{}) (interface{}, error) {
	req := reqest.(*pb.PalindromeRequest)
	return endpoints.PalindromeReq{Word: req.Word}, nil
}

func encodePalindromeResponse(_ context.Context, request interface{}) (interface{}, error) {
	resp := request.(endpoints.PalindromeResp)
	return &pb.PalindromeResponse{IsPalindrome: resp.Result}, nil
}
