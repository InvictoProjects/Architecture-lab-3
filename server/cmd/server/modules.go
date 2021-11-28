//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/invictoprojects/architecture-lab-3/server/balancers"
	"github.com/invictoprojects/architecture-lab-3/server/machines"
)

func ComposeApiServer(port HttpPortNumber) (*BalancersApiServer, error) {
	wire.Build(
		NewDbConnection,
		balancers.Providers,
		machines.Providers,
		wire.Struct(new(BalancersApiServer), "Port", "BalancersHandler", "MachinesHandler"),
	)
	return nil, nil
}
