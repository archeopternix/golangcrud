package view

import (
	"github.com/labstack/echo/v4"
)

func (s Server) handleDashboard(c echo.Context) error {
	p := NewPage("Dashboard", "Dashboard")
	return c.Render(http.StatusOK, "dashboard", p)
}
