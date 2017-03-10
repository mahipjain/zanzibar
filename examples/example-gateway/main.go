// Code generated by zanzibar
// @generated

package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/clients"
	"github.com/uber/zanzibar/examples/example-gateway/endpoints"
	"github.com/uber/zanzibar/runtime"
)

func getDirName() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}

func getConfigDirName() string {
	return filepath.Join(
		getDirName(),
		".",
		"config",
	)
}

func main() {
	config := zanzibar.NewStaticConfigOrDie([]string{
		filepath.Join(getDirName(), "zanzibar-defaults.json"),
		filepath.Join(getConfigDirName(), "production.json"),
		filepath.Join(os.Getenv("CONFIG_DIR"), "production.json"),
	}, nil)

	clients := clients.CreateClients(config)

	server, err := zanzibar.CreateGateway(config, &zanzibar.Options{
		Clients: clients,
	})
	if err != nil {
		panic(err)
	}

	err = server.Bootstrap(endpoints.Register)
	if err != nil {
		panic(err)
	}

	server.Logger.Info("Started ExampleGateway",
		zap.String("realAddr", server.RealAddr),
		zap.Object("config", config.InspectOrDie()),
	)

	// TODO: handle sigterm gracefully

	server.Wait()

	// TODO: emit metrics about startup.
	// TODO: setup and configure tracing/jeager.
}
