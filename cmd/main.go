package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/xvbnm48/-learning-go-kit/endpoints"
	"github.com/xvbnm48/-learning-go-kit/pb"
	"github.com/xvbnm48/-learning-go-kit/service"
	"github.com/xvbnm48/-learning-go-kit/transports"
	"google.golang.org/grpc"
)

func main() {
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	addService := service.NewService(logger)
	addEnpoint := endpoints.MakeEndpoints(addService)
	grpcServer := transports.NewGRPCServer(addEnpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":4848")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		pb.RegisterMathServiceServer(baseServer, grpcServer)
		// fmt.Printf("listening on port %s", grpcListener)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
