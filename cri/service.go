package cri

import (
	"fmt"

	"github.com/jeasonstudio/wasmer-cri/cri/config"
	"github.com/jeasonstudio/wasmer-cri/cri/utils"
	"google.golang.org/grpc"
)

// Service serves the kubelet runtime grpc api which will be consumed by kubelet.
type Service struct {
	config *config.Config
	server *grpc.Server
}

func RunCRIService(cfg *config.Config) (err error) {
	fmt.Println("Start CRI service with CRI version: v1alpha2")
	var service *Service

	service, err = NewService(cfg)

	if err != nil {
		return err
	}

	return service.Serve()
}

// NewService creates a brand new cri service.
func NewService(cfg *config.Config) (*Service, error) {
	s := &Service{
		config: cfg,
		server: grpc.NewServer(),
	}

	// runtime.RegisterRuntimeServiceServer(s.server, criMgr)
	// runtime.RegisterImageServiceServer(s.server, criMgr)
	// runtime.RegisterVolumeServiceServer(s.server, criMgr)

	// // EnableHandlingTimeHistogram turns on recording of handling time
	// // of RPCs. Histogram metrics can be very expensive for Prometheus
	// // to retain and query.
	// metrics.GRPCMetrics.EnableHandlingTimeHistogram()
	// // Initialize all metrics.
	// metrics.GRPCMetrics.InitializeMetrics(s.server)

	return s, nil
}

// Serve starts grpc server.
func (s *Service) Serve() error {
	listener, err := utils.GetListener(s.config.Listen, nil)
	if err != nil {
		return err
	}
	return s.server.Serve(listener)
}
