package provider

import (
	"time"

	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/config"
)

func NewTimeoutContextDuration(config config.Config) time.Duration {
	return config.Context.Timeout
}
