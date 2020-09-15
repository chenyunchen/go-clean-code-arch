//+build wireinject

//The build tag makes sure the stub is not built in the final build.
package http

import (
	"github.com/google/wire"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/provider"
)

func Initialize(configPath string) (*Application, error) {
	wire.Build(
		newApplication,
		provider.UseCaseSet,
		provider.DeliverySet,
		provider.RepositorySet,
		provider.ConfigSet,
	)

	return &Application{}, nil
}
