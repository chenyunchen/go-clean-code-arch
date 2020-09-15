//+build wireinject

//The build tag makes sure the stub is not built in the final build.
package http

import (
	"github.com/google/wire"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/config"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/delivery/http/middleware"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/provider"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/repository/mysql"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/usecase"
)

func Initialize(configPath string) (*Application, error) {
	wire.Build(appSet)

	return &Application{}, nil
}

var appSet = wire.NewSet(
	newApplication,
	config.NewConfig,
	provider.NewEcho,
	middleware.NewMiddleware,
	mysql.NewMysqlArticleRepository,
	mysql.NewMysqlAuthorRepository,
	provider.NewDBConn,
	usecase.NewArticleUsecase,
	provider.NewTimeoutContextDuration,
)
