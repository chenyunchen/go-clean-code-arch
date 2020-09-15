package provider

import (
	"time"

	"github.com/google/wire"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/config"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/usecase"
)

var UseCaseSet = wire.NewSet(
	usecase.NewArticleUsecase,
	NewTimeoutContextDuration,
)

func NewTimeoutContextDuration(config config.Config) time.Duration {
	return config.Context.Timeout
}
