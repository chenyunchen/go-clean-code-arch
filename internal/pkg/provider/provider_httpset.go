package provider

import (
	"github.com/google/wire"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/delivery/http/middleware"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/repository/mysql"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/usecase"
)

var HttpSet = wire.NewSet(
	NewEcho,
	middleware.NewMiddleware,
	NewTimeoutContextDuration,
	mysql.NewMysqlArticleRepository,
	mysql.NewMysqlAuthorRepository,
	usecase.NewArticleUsecase,
)
