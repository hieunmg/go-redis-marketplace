package chat

import (
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"

	"go-redis-marketplace/pkg/common"

	"go-redis-marketplace/pkg/config"
	"go-redis-marketplace/pkg/transport"
)

type GrpcServer struct {
	grpcPort string
	logger   common.GrpcLog
	s        *grpc.Server
	// userSvc  UserService
	// chanSvc  ChannelService
}

func NewGrpcServer(name string, logger common.GrpcLog, config *config.Config) *GrpcServer {
	srv := &GrpcServer{
		grpcPort: config.Chat.Grpc.Server.Port,
		logger:   logger,
		// userSvc:  userSvc,
		// chanSvc:  chanSvc,
	}
	srv.s = transport.InitializeGrpcServer(name, srv.logger)
	return srv
}

func (srv *GrpcServer) Register() {
	// chatpb.RegisterChannelServiceServer(srv.s, srv)
	// chatpb.RegisterUserServiceServer(srv.s, srv)
}

func (srv *GrpcServer) Run() {
	go func() {
		addr := "0.0.0.0:" + srv.grpcPort
		srv.logger.Info("grpc server listening", slog.String("addr", addr))
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			srv.logger.Error(err.Error())
			os.Exit(1)
		}
		if err := srv.s.Serve(lis); err != nil {
			srv.logger.Error(err.Error())
			os.Exit(1)
		}
	}()
}

func (srv *GrpcServer) GracefulStop() error {
	srv.s.GracefulStop()
	return nil
}

var UserConn *UserClientConn

type UserClientConn struct {
	Conn *grpc.ClientConn
}

func NewUserClientConn(config *config.Config) (*UserClientConn, error) {
	conn, err := transport.InitializeGrpcClient(config.Chat.Grpc.Client.User.Endpoint)
	if err != nil {
		return nil, err
	}
	UserConn = &UserClientConn{
		Conn: conn,
	}
	return UserConn, nil
}

var ForwarderConn *ForwarderClientConn

type ForwarderClientConn struct {
	Conn *grpc.ClientConn
}

func NewForwarderClientConn(config *config.Config) (*ForwarderClientConn, error) {
	conn, err := transport.InitializeGrpcClient(config.Chat.Grpc.Client.Forwarder.Endpoint)
	if err != nil {
		return nil, err
	}
	ForwarderConn = &ForwarderClientConn{
		Conn: conn,
	}
	return ForwarderConn, nil
}
