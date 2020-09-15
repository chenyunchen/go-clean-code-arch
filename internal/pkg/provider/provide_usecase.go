package provider

import (
	"time"

	"github.com/google/wire"
	"github.com/spf13/viper"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/usecase"
)

var UseCaseSet = wire.NewSet(
	usecase.NewArticleUsecase,
	NewTimeoutContextDuration,
)

func NewTimeoutContextDuration() time.Duration {
	return time.Duration(viper.GetInt("context.timeout")) * time.Second
}
