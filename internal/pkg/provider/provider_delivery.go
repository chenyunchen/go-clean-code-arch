package provider

import (
	"github.com/labstack/echo"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/delivery/http"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/delivery/http/middleware"
	"gitlab.silkrode.com.tw/team_golang/kbc2/sample/internal/pkg/domain"
)

func NewEcho(m *middleware.GoMiddleware, us domain.ArticleUsecase) *echo.Echo {
	e := echo.New()
	e.Use(m.CORS)
	http.NewArticleHandler(e, us)
	return e
}
