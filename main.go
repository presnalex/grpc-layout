package main

import (
	"context"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	pb "github.com/presnalex/grpc-layout/grpc-layout-proto/go/proto"
	"github.com/presnalex/grpc-layout/handler"
	jsoncodec "go.unistack.org/micro-codec-json/v3"
	consulconfig "go.unistack.org/micro-config-consul/v3"
	envconfig "go.unistack.org/micro-config-env/v3"
	sgrpc "go.unistack.org/micro-server-grpc/v3"

	"github.com/presnalex/go-micro/v3/service"
	"github.com/presnalex/statscheck"
	micro "go.unistack.org/micro/v3"
	"go.unistack.org/micro/v3/config"
	"go.unistack.org/micro/v3/logger"
)

const appName = "grpc-layout"

var (
	BuildDate  string
	AppVersion string
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-ch
		logger.Infof(ctx, "handle signal %v, exiting", sig)
		cancel()
	}()

	cfg := newConfig(appName, AppVersion)
	consulcfg := consulconfig.NewConfig(
		config.Struct(cfg),
		config.Codec(jsoncodec.NewCodec()),
		config.BeforeLoad(func(ctx context.Context, c config.Config) error {
			if len(cfg.Consul.NamespacePath) == 0 {
				cfg.Consul.NamespacePath = "go-micro-layouts"
			}
			if len(cfg.Consul.AppPath) == 0 {
				cfg.Consul.AppPath = "grpc-layout"
			}
			logger.Infof(ctx, "Consul Address: %s", cfg.Consul.Addr)
			logger.Infof(ctx, "Consul Path: %s", filepath.Join(cfg.Consul.NamespacePath, cfg.Consul.AppPath))
			return c.Init(
				consulconfig.Address(cfg.Consul.Addr),
				consulconfig.Token(cfg.Consul.Token),
				consulconfig.Path(filepath.Join(cfg.Consul.NamespacePath, cfg.Consul.AppPath)),
			)
		}),
	)
	if err := config.DefaultBeforeLoad(ctx, consulcfg); err != nil {
		logger.Fatalf(ctx, "failed to load config: %v", err)
	}

	err := config.Load(ctx,
		[]config.Config{
			config.NewConfig(
				config.Struct(cfg),
			),
			envconfig.NewConfig(
				config.Struct(cfg),
			),
			consulcfg},
	)
	if err != nil {
		logger.Fatalf(ctx, "failed to load config: %v", err)
	}

	serverConfig := &service.ServerConfig{
		Name:    cfg.Server.Name,
		Version: cfg.Server.Version,
		ID:      cfg.Server.ID,
		Addr:    cfg.Server.Addr,
	}
	srvOpts, _ := service.ServerOptions(serverConfig)
	srv := sgrpc.NewServer(srvOpts...)

	svc := micro.NewService(
		micro.Context(ctx),
		micro.Server(srv),
	)

	if err = svc.Init(); err != nil {
		logger.Fatal(ctx, err)
	}

	h := &handler.Handler{}

	err = svc.Init()
	if err != nil {
		logger.Fatal(ctx, err)
	}

	err = pb.RegisterPingPongServiceHandler(svc.Server(), h)
	if err != nil {
		logger.Fatalf(ctx, "can't register service handler: %v", err)
	}

	statsOpts := append([]statscheck.Option{},
		statscheck.WithDefaultHealth(),
		statscheck.WithMetrics(),
		statscheck.WithVersionDate(AppVersion, BuildDate),
	)

	if cfg.Core.Profile {
		statsOpts = append(statsOpts, statscheck.WithProfile())
	}

	healthServer := statscheck.NewServer(statsOpts...)
	go func() {
		logger.Fatal(ctx, healthServer.Serve(cfg.Metric.Addr))
	}()

	if err := svc.Run(); err != nil {
		logger.Fatal(ctx, err)
	}
}
