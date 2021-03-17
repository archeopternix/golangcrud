package view

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) getDashboard(c echo.Context) error {
	p := NewPage("Dashboard", "Dashboard")
	return c.Render(http.StatusOK, "dashboard", p)
}
