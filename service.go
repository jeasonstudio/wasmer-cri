package wasmercri

import (
	"net"
	"os"
	"strings"
	"syscall"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// Config of service
type Config struct {
	Network string // unix
	Address string // /tmp/wasmshim.sock
}

// Service to provide grpc server
type Service struct {
	Config     *Config
	GRPCServer *grpc.Server
}

func init() {
	logLevel := log.InfoLevel
	envLogLevel := os.Getenv("LOGLEVEL")

	switch strings.ToUpper(envLogLevel) {
	case "TRACE":
		logLevel = log.TraceLevel
		break
	case "DEBUG":
		logLevel = log.DebugLevel
		break
	case "INFO":
		logLevel = log.InfoLevel
		break
	case "WARN":
		logLevel = log.WarnLevel
		break
	case "ERROR":
		logLevel = log.ErrorLevel
		break
	case "FATAL":
		logLevel = log.FatalLevel
		break
	default:
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

// NewService create service
func NewService(config *Config) (*Service, error) {
	server := grpc.NewServer()
	imageServer, _ := NewImageServer()
	runtimeServer, _ := NewRuntimeServer()

	pb.RegisterImageServiceServer(server, imageServer)
	pb.RegisterRuntimeServiceServer(server, runtimeServer)
	// TODO: other service register

	// always remove the named socket from the fs if its there
	err := syscall.Unlink(config.Address)
	if err != nil {
		// not really important if it fails, so do nothing
		log.WithFields(log.Fields{
			"network": config.Network,
			"address": config.Address,
		}).Warnf("Failed to unlink/clear unix socket, never mind: %v", err)
	}

	return &Service{
		Config:     config,
		GRPCServer: server,
	}, nil
}

// Listen start unix socket
func (s *Service) Listen() error {
	listener, err := net.Listen(s.Config.Network, s.Config.Address)
	if err != nil {
		// not really important if it fails
		log.WithError(err).Fatal("Failed to `net.listen` when starting server.")
		return err
	}

	// Unix sockets must be unlink()ed before being reused again.
	// Unfortunately, this defer is not run when a signal is received, e.g. CTRL-C.
	defer func() {
		listener.Close()
	}()

	log.WithFields(log.Fields{
		"network": s.Config.Network,
		"address": s.Config.Address,
	}).Infof("Serving on %s://%s\n", s.Config.Network, listener.Addr().String())

	if err := s.GRPCServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	return nil
}
