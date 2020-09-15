package provider

import (
	"github.com/google/wire"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/config"
)

var ConfigSet = wire.NewSet(
	config.NewConfig,
)
